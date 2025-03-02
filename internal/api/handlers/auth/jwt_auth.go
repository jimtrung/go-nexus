package auth

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

    hashedPassword, err := services.HashPassword(req.Password)
    if err != nil {
        c.String(http.StatusInternalServerError, err.Error())
        return
    }
    req.Password = hashedPassword

    if err := services.InsertIntoUsers(req); err != nil {
        c.String(http.StatusInternalServerError, err.Error())
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
        c.String(http.StatusInternalServerError, "Wrong username/password")
        return
    }

    token, err := services.CreateSignedToken(req.Username)
    if err != nil {
        c.String(http.StatusInternalServerError, "Failed to sign token")
        return
    }

    c.SetSameSite(http.SameSiteLaxMode)
    c.SetCookie("Authorization", token, 3600 * 4, "/", "", true, true)

    c.Header("HX-Location", "/p/user/profile")
    c.Status(http.StatusOK)
}

func Validate(c *gin.Context) {
    usernameAny, _ := c.Get("username")
    username, ok := usernameAny.(string)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to get username",
        })
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Hello " + username,
    })
}

func Logout(c *gin.Context) {
    c.SetSameSite(http.SameSiteLaxMode)
    c.SetCookie("Authorization", "", 0, "/", "", true, true)

    c.Header("HX-Location", "/p/user/login")
    c.Status(http.StatusOK)
}
