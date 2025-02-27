package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/go-nexus/internal/domain/models"
	"github.com/jimtrung/go-nexus/internal/infra/db"
	"github.com/jimtrung/go-nexus/internal/services"
)

func Signup(c *gin.Context) {
	username, email, password := c.PostForm("username"), c.PostForm("email"), c.PostForm("password")
	if username == "" || password == "" || email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Empty field",
		})
		return
	}

    // Hash the password
    hashedPassword, err := services.HashPassword(password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
        return
    }

	req := models.User{
		Username: username,
		Email:    email,
		Password: hashedPassword,
	}

    if err := services.InsertIntoUsers(db.DB, req); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": err.Error(),
        })
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

	req := &models.LoginRequest{
		Username: username,
		Password: password,
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

	c.JSON(http.StatusOK, req)
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
