package pages

import (
	"github.com/gin-gonic/gin"
	"github.com/jimtrung/go-nexus/internal/api/handlers"
	"github.com/jimtrung/go-nexus/internal/infra/logger/zap"
	"github.com/jimtrung/go-nexus/templates/components/user"
)

func RenderSignupPage(c *gin.Context) {
    if err := handlers.Render(c, user.Signup()); err != nil {
        zap.NewLogger().Error("error", err.Error())
    }
}
