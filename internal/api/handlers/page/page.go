package page

import (
	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/jimtrung/go-nexus/internal/infra/logger/zap"
	"github.com/jimtrung/go-nexus/templates/component"
)

type PageHandler struct {
	Logger *zap.Logger
}

func NewPageLogger(logger *zap.Logger) *PageHandler {
	return &PageHandler{
		Logger: logger,
	}
}

func Render(c *gin.Context, component templ.Component) error {
	c.Header("Content-Type", "text/html")
	return component.Render(c.Request.Context(), c.Writer)
}

func (h *PageHandler) RenderHomePage(c *gin.Context) {
	if err := Render(c, component.Home()); err != nil {
		h.Logger.Error(err.Error())
	}
}

func (h *PageHandler) RenderLoginPage(c *gin.Context) {
	if err := Render(c, component.Login()); err != nil {
		h.Logger.Error(err.Error())
	}
}

func (h *PageHandler) RenderSignupPage(c *gin.Context) {
	if err := Render(c, component.Signup()); err != nil {
		h.Logger.Error(err.Error())
	}
}

func (h *PageHandler) RenderProfilePage(c *gin.Context) {
	// TODO: Add authentication middleware check
	component.Profile().Render(c.Request.Context(), c.Writer)
}
