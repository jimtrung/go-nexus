package authhandlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/go-nexus/internal/domain"
	"github.com/jimtrung/go-nexus/internal/infra/logger/zap"
	"github.com/jimtrung/go-nexus/internal/services"
)

type AuthHandler struct {
	AuthService *services.AuthService
	Logger      *zap.Logger
}

func NewAuthHandler(authServices *services.AuthService, logger *zap.Logger) *AuthHandler {
	return &AuthHandler{
		AuthService: authServices,
		Logger:      logger,
	}
}

func (h *AuthHandler) Signup(c *gin.Context) {
	req := &domain.User{}

	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong JSON format",
		})
		h.Logger.Error("Wrong JSON format", err.Error())
		return
	}

	if err := h.AuthService.SignUp(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error(err.Error())
		return
	}

	h.Logger.Info(fmt.Sprintf("User %s signup successfully", req.Username))
	c.JSON(http.StatusOK, gin.H{"name": req.Username})
}

func (h *AuthHandler) Login(c *gin.Context) {
	req := &domain.User{}

	if err := c.Bind(&req); err != nil {
		c.String(http.StatusBadRequest, "Wrong JSON format")
		h.Logger.Error("Wrong JSON format", err.Error())
		return
	}

	token, err := h.AuthService.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error(err.Error())
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", token, 3600*4, "/", "", true, true)
	h.Logger.Info(fmt.Sprintf("User %s login successfully", req.Username))

	c.Header("HX-Redirect", "/profile")
	c.Status(http.StatusOK)
}

func (h *AuthHandler) Logout(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", "", -1, "/", "", true, true)
	h.Logger.Info("User logged out successfully")

	c.Header("HX-Redirect", "/login")
	c.Status(http.StatusOK)
}

func (h *AuthHandler) ForgotPassword(c *gin.Context) {
	req := &domain.User{}
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error("Wrong JSON format", err.Error())
		return
	}

	if err := h.AuthService.ForgotPassword(req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error(err.Error())
		return
	}

	h.Logger.Info("Reset password sent successfully")
	c.JSON(http.StatusOK, gin.H{
		"message": "Email sent",
	})
}

func (h *AuthHandler) ResetPassword(c *gin.Context) {
	var req domain.ResetPassword
	if err := c.Bind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong JSON format",
		})
		h.Logger.Error("Wrong JSON format", err.Error())
		return
	}

	if req.Password != req.ConfirmPassword {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Password do not match",
		})
		h.Logger.Error("Password do not match")
		return
	}

	hash, err := services.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error(err.Error())
		return
	}

	if err := h.AuthService.AuthRepo.UpdatePassword(req.Token, hash); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error(err.Error())
		return
	}

	h.Logger.Info("Reset password successfully")
	c.JSON(http.StatusOK, gin.H{
		"message": "Reset password successfully",
	})
}

func (h *AuthHandler) Verify(c *gin.Context) {
	token := c.Param("token")
	if token == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty token",
		})
		h.Logger.Error("Empty token")
		return
	}

	if err := h.AuthService.AuthRepo.Verify(token); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		h.Logger.Error(err.Error())
		return
	}

	h.Logger.Info("User verified")
	c.JSON(http.StatusOK, gin.H{
		"message": "User verified",
	})
}
