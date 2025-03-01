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
    }

    hashedPassword, err := services.HashPassword(req.Password)
    if err != nil {
        c.String(http.StatusInternalServerError, err.Error())
    }
    req.Password = hashedPassword

    if err := services.InsertIntoUsers(req); err != nil {
        c.String(http.StatusInternalServerError, err.Error())
        return
    }

	c.JSON(http.StatusOK, gin.H{
        "message": "User added succesfully",
    })
}

func Login(c *gin.Context) {
	username, password := c.PostForm("username"), c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty field",
		})
		return
	}

	req := models.User{
		Username: username,
		Password: password,
	}

    // Compare password
    if err := services.IsValidUser(req); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

    token, err := services.CreateSignedToken(username)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to sign token",
        })
        return
    }

    c.SetSameSite(http.SameSiteLaxMode)
    c.SetCookie("Authorization", token, 3600 * 4, "/", "", false, true)
	c.JSON(http.StatusOK, gin.H{
        "message": "Login successfully",
    })
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
