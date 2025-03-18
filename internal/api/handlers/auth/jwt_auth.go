package authhandlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/go-nexus/internal/domain"
	"github.com/jimtrung/go-nexus/internal/infra/logger/zap"
	"github.com/jimtrung/go-nexus/internal/services"
)

var logger = zap.NewLogger()

func Signup(c *gin.Context) {
	var req domain.User

	if err := c.Bind(&req); err != nil {
		c.String(http.StatusBadRequest, "Wrong JSON format")
        logger.Error("Wrong JSON format", err.Error())
		return
	}

	if !services.IsValidEmail(req.Email) || !services.HasMXRecord(req.Email) {
		c.String(http.StatusBadRequest, "Invalid email")
        logger.Error("Invalid email")
		return
	}

	hashedPassword, err := services.HashPassword(req.Password)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
        logger.Error(err.Error())
		return
	}

	req.Password = hashedPassword
	token := services.GenerateToken()
    req.Token = token

	if err := services.InsertIntoUsers(req); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
        logger.Error(err.Error())
		return
	}

	err = services.SendVerificationEmail(req.Email, token)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
        logger.Error(err.Error())
		return
	}

    logger.Info(fmt.Sprintf("User %s signup successfully", req.Username))

	c.Header("HX-Location", "/p/user/login")
	c.Status(http.StatusOK)
}

func Login(c *gin.Context) {
	var req domain.User

	if err := c.Bind(&req); err != nil {
		c.String(http.StatusBadRequest, "Wrong JSON format")
        logger.Error("Wrong JSON format", err.Error())
		return
	}

	if err := services.IsValidUser(req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
        logger.Error(err.Error())
		return
	}

	token, err := services.CreateSignedToken(req.Username)
	if err != nil {
		c.String(http.StatusInternalServerError, err.Error())
        logger.Error(err.Error())
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*4, "/", "", true, true)
	c.Header("HX-Location", "/p/user/profile")
	c.Status(http.StatusOK)
    logger.Info(fmt.Sprintf("User %s login successfully", req.Username))
}

func Logout(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", "", 0, "/", "", true, true)
    c.Redirect(http.StatusMovedPermanently, "/p/user/login")
}

func ForgotPassword(c *gin.Context) {
    var userReq domain.User

    if err := c.Bind(&userReq); err != nil {
        c.String(http.StatusBadRequest, "Wrong JSON format")
        logger.Error("Wrong JSON format", err.Error())
        return
    }

	if !services.IsValidEmail(userReq.Email) || !services.HasMXRecord(userReq.Email) {
		c.String(http.StatusBadRequest, "Invalid email")
        logger.Error("Invalid email")
		return
	}

    token := services.GenerateToken()
    userReq.Token = token

    if err := services.AddTokenToUser(userReq.Email, userReq.Token); err != nil {
        c.String(http.StatusBadRequest, "Failed to add token to user")
        logger.Error("Failed to add token to user", err.Error())
        return
    }

    if err := services.ResetPasswordEmail(userReq.Email, userReq.Token); err != nil {
        c.String(http.StatusBadRequest, "Failed to send reset email to user")
        logger.Error("Failed to send reset email to user")
        return
    }

    go func() {
        time.Sleep(time.Second * 300)
        services.RemoveToken(userReq.Token)
        logger.Info(fmt.Sprintf("Token removed from user %s", userReq.Username))
    }()

    c.Status(http.StatusOK)
    logger.Info("Reset password sent successfully")
}

func ResetPassword(c *gin.Context) {
    var resetReq domain.ResetPassword

    if err := c.Bind(&resetReq); err != nil {
        c.String(http.StatusBadRequest, "Wrong JSON format")
        logger.Error("Wrong JSON format", err.Error())
        return
    }

    if resetReq.Password != resetReq.ConfirmPassword {
        c.String(http.StatusInternalServerError, "Password do not match")
        logger.Error("Password do not match")
        return
    }

    if err := services.ResetPassword(resetReq.Token, resetReq.Password); err != nil {
        c.String(http.StatusInternalServerError, err.Error())
        logger.Error(err.Error())
        return
    }

	c.Header("HX-Location", "/p/user/login")
	c.Status(http.StatusOK)
    logger.Info("Reset password successfully")
}
