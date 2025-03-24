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
	Logger       *zap.Logger
	AuthService  *services.AuthService
	FriendService *services.FriendService
}

func NewPageHandler(logger *zap.Logger, authService *services.AuthService, friendService *services.FriendService) *PageHandler {
	return &PageHandler{
		Logger:       logger,
		AuthService:  authService,
		FriendService: friendService,
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
		c.Redirect(http.StatusSeeOther, "/p/login")
		return
	}

	userID := uint(userIDInt.(int))
	user, err := h.AuthService.GetUserByID(userID)
	if err != nil {
		h.Logger.Error("Failed to get user info", err.Error())
		c.Redirect(http.StatusSeeOther, "/p/login")
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
		c.Redirect(http.StatusSeeOther, "/p/login")
		return
	}

	if err := Render(c, component.ResetPassword(component.ResetPasswordProps{Token: token})); err != nil {
		h.Logger.Error(err.Error())
	}
}

func (h *PageHandler) RenderFriendsPage(c *gin.Context) {
	userIDInt, exists := c.Get("user_id")
	if !exists {
		c.Redirect(http.StatusSeeOther, "/p/login")
		return
	}

	userID := uint(userIDInt.(int))

	friends, err := h.FriendService.GetAllFriends(userID)
	if err != nil {
		h.Logger.Error("Failed to get friends", err.Error())
		c.Redirect(http.StatusSeeOther, "/p/login")
		return
	}

	pendingRequests, err := h.FriendService.GetPendingRequests(userID)
	if err != nil {
		h.Logger.Error("Failed to get pending requests", err.Error())
		c.Redirect(http.StatusSeeOther, "/p/login")
		return
	}

	incomingRequests, err := h.FriendService.GetPendingRequests(userID)
	if err != nil {
		h.Logger.Error("Failed to get incoming requests", err.Error())
		c.Redirect(http.StatusSeeOther, "/p/login")
		return
	}

	friendsProps := component.FriendsProps{
		Friends:          make([]component.Friend, len(friends)),
		PendingRequests:  make([]component.Friend, len(pendingRequests)),
		IncomingRequests: make([]component.Friend, len(incomingRequests)),
	}

	for i, friend := range friends {
		friendsProps.Friends[i] = component.Friend{
			FriendID:   friend.FriendID,
			SenderID:   friend.SenderID,
			ReceiverID: friend.ReceiverID,
			Status:     string(friend.Status),
			CreatedAt:  friend.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:  friend.UpdatedAt.Format("2006-01-02 15:04:05"),
			Username:   "User", // TODO: Get username from user service
			Avatar:     "https://ui-avatars.com/api/?name=User&background=random",
		}
	}

	for i, request := range pendingRequests {
		friendsProps.PendingRequests[i] = component.Friend{
			FriendID:   request.FriendID,
			SenderID:   request.SenderID,
			ReceiverID: request.ReceiverID,
			Status:     string(request.Status),
			CreatedAt:  request.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:  request.UpdatedAt.Format("2006-01-02 15:04:05"),
			Username:   "User", // TODO: Get username from user service
			Avatar:     "https://ui-avatars.com/api/?name=User&background=random",
		}
	}

	for i, request := range incomingRequests {
		friendsProps.IncomingRequests[i] = component.Friend{
			FriendID:   request.FriendID,
			SenderID:   request.SenderID,
			ReceiverID: request.ReceiverID,
			Status:     string(request.Status),
			CreatedAt:  request.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:  request.UpdatedAt.Format("2006-01-02 15:04:05"),
			Username:   "User", // TODO: Get username from user service
			Avatar:     "https://ui-avatars.com/api/?name=User&background=random",
		}
	}

	if err := Render(c, component.Friends(friendsProps)); err != nil {
		h.Logger.Error(err.Error())
	}
}
