package pageshandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/go-nexus/internal/api/handlers"
	"github.com/jimtrung/go-nexus/internal/domain/models"
	"github.com/jimtrung/go-nexus/internal/infra/logger/zap"
	"github.com/jimtrung/go-nexus/internal/services"
	"github.com/jimtrung/go-nexus/templates/components"
	userComponents "github.com/jimtrung/go-nexus/templates/components/user"
)

var logger = zap.NewLogger()

func RenderSignupPage(c *gin.Context) {
    if err := handlers.Render(c, userComponents.Signup()); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to render the page",
        })
        logger.Error("Failed to render the page", err.Error())
        return
    }
    logger.Info("Signup page rendered successfully")
}

func RenderLoginPage(c *gin.Context) {
    if err := handlers.Render(c, userComponents.Login()); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to render the page",
        })
        logger.Error("Failed to render the page", err.Error())
        return
    }
    logger.Info("Login page rendered successfully")
}

func RenderProfilePage(c *gin.Context) {
    data, exists := c.Get("username")
    if !exists {
        logger.Error("Username has not been set")
        return
    }

    username := data.(string)
    if username != data.(string) {
        logger.Error("Wrong data type")
        return
    }

    userData, err := services.GetUserByUsername(username)
    if err != nil {
        logger.Error("Failed to get user", err.Error())
        return
    }

    if err := handlers.Render(
        c, userComponents.ProfilePage(
            userData, userComponents.ProfileContent(userData)),
        ); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to render the page",
        })
        logger.Error("Failed to render the page", err.Error())
        return
    }
    logger.Info("Profile page rendered successfully")
}

func RenderForgotPasswordPage(c *gin.Context) {
    if err := handlers.Render(c, userComponents.ForgotPassword()); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to render the page",
        })
        logger.Error("Failed to render the page", err.Error())
        return
    }
    logger.Info("Forgot password page rendered successfully")
}

func RenderVerifyPage(c *gin.Context) {
    if err := handlers.Render(c, userComponents.Verify()); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to render the page",
        })
        logger.Error("Failed to render the page", err.Error())
        return
    }

    token := c.Param("token")
    if err := services.VerifyUser(token); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H {
            "error": err.Error(),
        })
        return
    }
    logger.Info("Verify page rendered successfully")
}

func RenderResetPasswordPage(c *gin.Context) {
    if err := handlers.Render(c, userComponents.ResetPassword()); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to render the page",
        })
        logger.Error("Failed to render the page", err.Error())
        return
    }
    logger.Info("Verify page rendered successfully")
}

func RenderEditProfilePage(c *gin.Context) {
    if err := handlers.Render(
        c, userComponents.ProfilePage(
            models.User{}, userComponents.EditProfileContent(models.User{})),
        ); err != nil {
        logger.Error("Failed to render the page", err.Error())
        return
    }
    logger.Info("Edit profile page rendered successfully")
}

func RenderSecurityPage(c *gin.Context) {
    if err := handlers.Render(
        c, userComponents.ProfilePage(
            models.User{}, userComponents.SecurityContent()),
        ); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to render the page",
        })
        logger.Error("Failed to render the page", err.Error())
        return
    }
    logger.Info("Security page rendered successfully")
}

func RenderPreferencesPage(c *gin.Context) {
    if err := handlers.Render(
        c, userComponents.ProfilePage(
            models.User{}, components.PreferencesContent()),
        ); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to render the page",
        })
        logger.Error("Failed to render the page", err.Error())
        return
    }
    logger.Info("Preferences page rendered successfully")
}
