package routes

import (
	"github.com/jimtrung/go-nexus/internal/api/handlers/pages"
	"github.com/jimtrung/go-nexus/internal/middleware"
)

func SetupPagesRoutes(r *Routes) {
    p := r.Router.Group("/p")
    {
        user := p.Group("/user")
        {
            user.GET("/signup", pages.RenderSignupPage)
            user.GET("/login", pages.RenderLoginPage)
            user.GET("/profile", middleware.RequireAuth, pages.RenderProfilePage)
            user.GET("/forgot-password", pages.RenderForgotPasswordPage)
        }
    }

    r.Router.Static("/static", "./static")
}
