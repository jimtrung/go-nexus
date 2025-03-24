package friend

import (
	"github.com/gin-gonic/gin"
	"github.com/jimtrung/go-nexus/internal/infra/logger/zap"
	"github.com/jimtrung/go-nexus/internal/services"
	"net/http"
	"github.com/jimtrung/go-nexus/internal/domain"
	"strconv"
)

type FriendHandler struct {
	FriendService *services.FriendService
	Logger        *zap.Logger
}

func NewFriendHandler(friendService *services.FriendService, logger *zap.Logger) *FriendHandler {
	return &FriendHandler{
		FriendService: friendService,
		Logger:        logger,
	}
}

func (h *FriendHandler) CreateRequest(c *gin.Context) {
	userID := c.GetUint("user_id")

	req := &domain.Friend{}
	if err := c.BindJSON(req); err != nil {
		h.Logger.Error("Failed to bind request", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong JSON format",
		})
		return
	}
	req.SenderID = userID

	if err := h.FriendService.CreateRequest(req); err != nil {
		h.Logger.Error("Failed to create request", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to create request",
		})
		return
	}

	h.Logger.Info("Request created successfully")
	c.JSON(http.StatusCreated, gin.H{
		"message": "Request created successfully",
	})
}

func (h *FriendHandler) GetAllFriends(c *gin.Context) {
	userID := c.GetUint("user_id")

	friends, err := h.FriendService.GetAllFriends(userID)
	if err != nil {
		h.Logger.Error("Failed to get all friends", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get all friends",
		})
		return
	}

	h.Logger.Info("Friends fetched successfully")
	c.JSON(http.StatusOK, gin.H{
		"friends": friends,
	})
}

func (h *FriendHandler) GetPendingRequests(c *gin.Context) {
	userID := c.GetUint("user_id")

	requests, err := h.FriendService.GetPendingRequests(userID)
	if err != nil {
		h.Logger.Error("Failed to get pending requests", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get pending requests",
		})
		return
	}

	h.Logger.Info("Pending requests fetched successfully")
	c.JSON(http.StatusOK, gin.H{
		"requests": requests,
	})
}

func (h *FriendHandler) AcceptRequest(c *gin.Context) {
	userID := c.GetUint("user_id")

	req := &domain.Friend{}
	if err := c.BindJSON(req); err != nil {
		h.Logger.Error("Failed to bind request", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong JSON format",
		})
		return
	}
	req.ReceiverID = userID
	
	if err := h.FriendService.AcceptRequest(req); err != nil {
		h.Logger.Error("Failed to accept request", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to accept request",
		})
		return
	}

	h.Logger.Info("Request accepted successfully")
	c.JSON(http.StatusOK, gin.H{
		"message": "Request accepted successfully",
	})
}

func (h *FriendHandler) RejectRequest(c *gin.Context) {
	userID := c.GetUint("user_id")

	req := &domain.Friend{}
	if err := c.BindJSON(req); err != nil {
		h.Logger.Error("Failed to bind request", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Wrong JSON format",
		})
		return
	}
	req.ReceiverID = userID
	
	if err := h.FriendService.RejectRequest(req); err != nil {
		h.Logger.Error("Failed to reject request", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to reject request",
		})
		return
	}

	h.Logger.Info("Request rejected successfully")
	c.JSON(http.StatusOK, gin.H{
		"message": "Request rejected successfully",
	})
}

func (h *FriendHandler) CancelRequest(c *gin.Context) {
	userID := c.GetUint("user_id")
	receiverIDString := c.Param("receiver_id")
	receiverID64, err := strconv.ParseUint(receiverIDString, 10, 64)
	receiverID := uint(receiverID64)
	if err != nil {
		h.Logger.Error("Failed to parse receiver ID", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid receiver ID",
		})
		return
	}

	req := &domain.Friend{
		SenderID: userID,
		ReceiverID: receiverID,
	}

	if err := h.FriendService.CancelRequest(req); err != nil {
		h.Logger.Error("Failed to cancel request", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to cancel request",
		})
		return
	}	

	h.Logger.Info("Request canceled successfully")
	c.JSON(http.StatusOK, gin.H{
		"message": "Request canceled successfully",
	})
}

func (h *FriendHandler) RemoveFriend(c *gin.Context) {
	userID := c.GetUint("user_id")
	receiverIDString := c.Param("receiver_id")
	receiverID64, err := strconv.ParseUint(receiverIDString, 10, 64)
	receiverID := uint(receiverID64)
	if err != nil {
		h.Logger.Error("Failed to parse receiver ID", err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid receiver ID",
		})
		return
	}
	req := &domain.Friend{
		SenderID: userID,
		ReceiverID: receiverID,
	}
	
	if err := h.FriendService.RemoveFriend(req); err != nil {
		h.Logger.Error("Failed to remove friend", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to remove friend",
		})
		return
	}

	h.Logger.Info("Friend removed successfully")
	c.JSON(http.StatusOK, gin.H{
		"message": "Friend removed successfully",
	})
}