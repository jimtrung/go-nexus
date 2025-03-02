package auth

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"github.com/jimtrung/go-nexus/internal/services"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/facebook"
	"github.com/markbates/goth/providers/google"
)

const (
	key    = "12345678"
	maxAge = 86400 * 30
	isProd = true
)

func NewOAuth() {
	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(maxAge)
	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(
		google.New(
            os.Getenv("GOOGLE_CLIENT_ID"),
            os.Getenv("GOOGLE_CLIENT_SECRET"),
			"http://127.0.0.1:8080/auth/google/callback",
		),
        facebook.New(
            os.Getenv("FACEBOOK_CLIENT_ID"),
            os.Getenv("FACEBOOK_CLIENT_SECRET"),
            "http://127.0.0.1:8080/auth/facebook/callback",
        ),
	)
}

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
