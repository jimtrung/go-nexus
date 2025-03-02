package pages

import (
	"github.com/gin-gonic/gin"
	"github.com/jimtrung/go-nexus/internal/api/handlers"
	"github.com/jimtrung/go-nexus/internal/infra/logger/zap"
	"github.com/jimtrung/go-nexus/internal/services"
	"github.com/jimtrung/go-nexus/templates/components/user"
)

func RenderSignupPage(c *gin.Context) {
    if err := handlers.Render(c, user.Signup()); err != nil {
        zap.NewLogger().Error("error", err.Error())
    }
}

func RenderLoginPage(c *gin.Context) {
    if err := handlers.Render(c, user.Login()); err != nil {
        zap.NewLogger().Error("error", err.Error())
    }
}

func RenderProfilePage(c *gin.Context) {
    data, exists := c.Get("username")
    if !exists {
        zap.NewLogger().Error("error", "Username has not been set")
        return
    }
    username := data.(string)
    if username != data.(string) {
        zap.NewLogger().Error("error", "Wrong data")
        return
    }

    userData, err := services.GetUserByUsername(username)
    if err != nil {
        zap.NewLogger().Error("error", "Failed to get user")
        return
    }

    zap.NewLogger().Info("data", userData)
    if err := handlers.Render(c, user.Profile(userData)); err != nil {
        zap.NewLogger().Error("error", err.Error())
    }
}
