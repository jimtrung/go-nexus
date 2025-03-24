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
	pageHandler := page.NewPageLogger(logger, authService)

	r.Router.GET("/", pageHandler.RenderHomePage)
	r.Router.GET("/login", pageHandler.RenderLoginPage)
	r.Router.GET("/signup", pageHandler.RenderSignupPage)
	r.Router.GET("/profile", middleware.RequireAuth, pageHandler.RenderProfilePage)
	r.Router.GET("/forgot-password", pageHandler.RenderForgotPasswordPage)
	r.Router.GET("/reset-password/:token", pageHandler.RenderResetPasswordPage)
}
