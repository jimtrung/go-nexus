package routes

import (
	pagesHandlers "github.com/jimtrung/go-nexus/internal/api/handlers/pages"
	"github.com/jimtrung/go-nexus/internal/middleware"
)

func SetupPagesRoutes(r *Routes) {
    p := r.Router.Group("/p")
    {
        user := p.Group("/user")
        {
            user.GET("/signup", pagesHandlers.RenderSignupPage)
            user.GET("/login", pagesHandlers.RenderLoginPage)
            user.GET("/verify/:token", pagesHandlers.RenderVerifyPage)
            user.GET("/forgot-password", pagesHandlers.RenderForgotPasswordPage)
            user.GET("/reset-password/:token", pagesHandlers.RenderResetPasswordPage)
            user.GET("/profile", middleware.RequireAuth, pagesHandlers.RenderProfilePage)
            user.GET("/profile/edit", pagesHandlers.RenderEditProfilePage)
            user.GET("/security", pagesHandlers.RenderSecurityPage)
        }
        p.GET("/preferences", pagesHandlers.RenderPreferencesPage)

        friends := p.Group("/friends")
        {
            friends.GET("/show", pagesHandlers.RenderShowPage)
        }
    }

    r.Router.Static("/static", "./static")
}
