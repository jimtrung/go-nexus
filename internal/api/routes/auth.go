package routes

import (
	authHandlers "github.com/jimtrung/go-nexus/internal/api/handlers/auth"
)

func SetupAuthRoutes(r *Routes) {
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
