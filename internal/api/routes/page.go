package routes

import (
	"github.com/jimtrung/go-nexus/internal/api/handlers/page"
	"github.com/jimtrung/go-nexus/internal/infra/logger/zap"
	"github.com/jimtrung/go-nexus/internal/middleware"
	"github.com/jimtrung/go-nexus/internal/repository"
	"github.com/jimtrung/go-nexus/internal/services"
)

func (r *Routes) SetupPageRoutes(logger *zap.Logger) {
	r.Router.Static("/static", "./static")

	authRepo := repository.NewUserRepository(r.Conn)
	authService := services.NewAuthService(authRepo)
	friendRepo := repository.NewFriendRepository(r.Conn)
	friendService := services.NewFriendService(friendRepo)
	pageHandler := page.NewPageHandler(logger, authService, friendService)

	pageRouter := r.Router.Group("/p")
	{
		pageRouter.GET("/", pageHandler.RenderHomePage)
		pageRouter.GET("/login", pageHandler.RenderLoginPage)
		pageRouter.GET("/signup", pageHandler.RenderSignupPage)
		pageRouter.GET("/profile", middleware.RequireAuth, pageHandler.RenderProfilePage)
		pageRouter.GET("/forgot-password", pageHandler.RenderForgotPasswordPage)
		pageRouter.GET("/reset-password/:token", pageHandler.RenderResetPasswordPage)
		pageRouter.GET("/friends", middleware.RequireAuth, pageHandler.RenderFriendsPage)
	}
}