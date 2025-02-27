package routes

import (
	authHandlers "github.com/jimtrung/go-nexus/internal/api/handlers/auth"
	"github.com/jimtrung/go-nexus/internal/middlware"
)

func SetupAuthRoutes(r *Routes) {
    auth := r.Router.Group("/auth")
    {
        // Oauth with google, oauth with facebook
        auth.GET("/:provider")
        auth.GET("/:provider/callback")

        // Manually login with password + email
        auth.POST("/login", authHandlers.Login)
        auth.POST("/signup", authHandlers.Signup)

        auth.GET("/validate", middlware.RequireAuth, authHandlers.Validate)
    }
}
