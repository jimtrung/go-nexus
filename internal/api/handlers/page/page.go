package page

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/gin-gonic/gin"
	"github.com/jimtrung/go-nexus/internal/infra/logger/zap"
	"github.com/jimtrung/go-nexus/internal/services"
	"github.com/jimtrung/go-nexus/templates/component"
)

type PageHandler struct {
	Logger      *zap.Logger
	AuthService *services.AuthService
}

func NewPageLogger(logger *zap.Logger, authService *services.AuthService) *PageHandler {
	return &PageHandler{
		Logger:      logger,
		AuthService: authService,
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
	userIDInt, exists := c.Get("user_id")
	if !exists {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	userID := uint(userIDInt.(int))
	user, err := h.AuthService.GetUserByID(userID)
	if err != nil {
		h.Logger.Error("Failed to get user info", err.Error())
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	if err := Render(c, component.Profile(user)); err != nil {
		h.Logger.Error(err.Error())
	}
}

func (h *PageHandler) RenderForgotPasswordPage(c *gin.Context) {
	if err := Render(c, component.ForgotPassword()); err != nil {
		h.Logger.Error(err.Error())
	}
}

func (h *PageHandler) RenderResetPasswordPage(c *gin.Context) {
	token := c.Param("token")
	if token == "" {
		c.Redirect(http.StatusSeeOther, "/login")
		return
	}

	if err := Render(c, component.ResetPassword(component.ResetPasswordProps{Token: token})); err != nil {
		h.Logger.Error(err.Error())
	}
}
