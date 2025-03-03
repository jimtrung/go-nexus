package authhandlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/go-nexus/internal/domain/models"
	"github.com/jimtrung/go-nexus/internal/infra/logger/zap"
	"github.com/jimtrung/go-nexus/internal/services"
)

func Signup(c *gin.Context) {
	var req models.User
	if err := c.Bind(&req); err != nil {
		c.String(http.StatusInternalServerError, "Wrong JSON format")
		return
	}

	if !services.IsValidEmail(req.Email) || !services.HasMXRecord(req.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid email",
        })
		return
	}

	hashedPassword, err := services.HashPassword(req.Password)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	req.Password = hashedPassword
	token := services.GenerateToken()
    req.Token = token

	if err := services.InsertIntoUsers(req); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = services.SendVerificationEmail(req.Email, token)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
		return
	}

	c.Header("HX-Location", "/p/user/login")
	c.Status(http.StatusOK)
}

func Login(c *gin.Context) {
	var req models.User

	if err := c.Bind(&req); err != nil {
		c.String(http.StatusBadRequest, "Wrong JSON format")
		return
	}

	if err := services.IsValidUser(req); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	token, err := services.CreateSignedToken(req.Username)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*4, "/", "", true, true)

	c.Header("HX-Location", "/p/user/profile")
	c.Status(http.StatusOK)
}

func Logout(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", "", 0, "/", "", true, true)

    c.Redirect(http.StatusMovedPermanently, "/p/user/login")
}

func ForgotPassword(c *gin.Context) {
    var userReq models.User

    if err := c.Bind(&userReq); err != nil {
        c.String(http.StatusBadRequest, "Wrong JSON format")
        return
    }

	if !services.IsValidEmail(userReq.Email) || !services.HasMXRecord(userReq.Email) {
		c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid email",
        })
		return
	}

    token := services.GenerateToken()
    userReq.Token = token
    if err := services.AddTokenToUser(userReq.Email, userReq.Token); err != nil {
        c.String(http.StatusBadRequest, "Failed to add token to user")
        return
    }

    if err := services.ResetPasswordEmail(userReq.Email, userReq.Token); err != nil {
        c.String(http.StatusBadRequest, "Failed to send reset email to user")
        return
    }

    go func() {
        time.Sleep(time.Second * 300)
        services.RemoveToken(userReq.Token)
        zap.NewLogger().Info("message", "Token removed from user")
    }()

    c.JSON(http.StatusOK, gin.H{
        "email": userReq.Email,
    })
}

func ResetPassword(c *gin.Context) {
    var resetReq models.ResetPassword
    if err := c.Bind(&resetReq); err != nil {
        c.String(http.StatusBadRequest, "Wrong JSON format")
        return
    }

    token := c.PostForm("token")
    resetReq.Token = token

    if resetReq.Password != resetReq.ConfirmPassword {
        c.String(http.StatusInternalServerError, "Password do not match")
        return
    }

    if err := services.ResetPassword(resetReq.Token, resetReq.Password); err != nil {
        c.String(http.StatusInternalServerError, "Failed to reset the password")
        return
    }

	c.Header("HX-Location", "/p/user/login")
	c.Status(http.StatusOK)
}
