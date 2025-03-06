package authhandlers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/go-nexus/internal/infra/logger/zap"
	"github.com/jimtrung/go-nexus/internal/services"
	"github.com/markbates/goth/gothic"
)

func BeginAuthProviderCallback(c *gin.Context) {
	provider := c.Param("provider")
	c.Request = c.Request.WithContext(context.WithValue(
		context.Background(),
		"provider",
		provider,
	))
	gothic.BeginAuthHandler(c.Writer, c.Request)
}

func GetAuthCallBackFunction(c *gin.Context) {
	provider := c.Param("provider")
	c.Request = c.Request.WithContext(context.WithValue(
		context.Background(),
		"provider",
		provider,
	))

	user, err := gothic.CompleteUserAuth(c.Writer, c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

    userInfo, err := services.SignupIfNotExist(user.Email)
    if err != nil {
        zap.NewLogger().Error("error", err.Error())
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    signedToken, err := services.CreateSignedToken(userInfo.Username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to create a signed token",
        })
        return
    }

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", signedToken, 3600*24, "/", "", false, true)

	c.Redirect(http.StatusFound, "http://127.0.0.1:8080/p/user/profile")
}
