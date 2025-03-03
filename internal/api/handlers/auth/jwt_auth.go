package authhandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/go-nexus/internal/domain/models"
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

	c.Header("HX-Location", "/p/user/login")
	c.Status(http.StatusOK)
}
