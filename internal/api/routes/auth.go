package routes

import (
	authHandlers "github.com/jimtrung/go-nexus/internal/api/handlers/auth"
	"github.com/jimtrung/go-nexus/internal/infra/logger/zap"
	"github.com/jimtrung/go-nexus/internal/repository"
	"github.com/jimtrung/go-nexus/internal/services"
)

func (r *Routes) SetupAuthRoutes(logger *zap.Logger) {
    authRepo := repository.NewUserRepository(r.Conn)
    authServices := services.NewAuthService(authRepo)
    authHandlers := authHandlers.NewAuthHandler(authServices, logger)
    auth := r.Router.Group("/auth")
    {
        auth.GET("/:provider", authHandlers.BeginAuthProviderCallback)
        auth.GET("/:provider/callback", authHandlers.GetAuthCallBackFunction)

        auth.POST("/signup", authHandlers.Signup)
        auth.POST("/login", authHandlers.Login)
        auth.GET("/logout", authHandlers.Logout)
        auth.POST("/forgot-password", authHandlers.ForgotPassword)
        auth.POST("/reset-password", authHandlers.ResetPassword)
    }
}
