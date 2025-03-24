package routes

import (
	"github.com/jimtrung/go-nexus/internal/api/handlers/page"
	"github.com/jimtrung/go-nexus/internal/infra/logger/zap"
)

func (r *Routes) SetupPageRoutes(logger *zap.Logger) {
	r.Router.Static("/static", "./static")
	pageHandler := page.NewPageLogger(logger)

	p := r.Router.Group("/p")
	{
		p.GET("/", pageHandler.RenderHomePage)
		p.GET("/login", pageHandler.RenderLoginPage)
		p.GET("/signup", pageHandler.RenderSignupPage)
	}

	r.Router.GET("/", pageHandler.RenderHomePage)
	r.Router.GET("/login", pageHandler.RenderLoginPage)
	r.Router.GET("/signup", pageHandler.RenderSignupPage)
}
