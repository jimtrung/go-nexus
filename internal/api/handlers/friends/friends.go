package friends

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/go-nexus/internal/infra/logger/zap"
	"github.com/jimtrung/go-nexus/internal/services"
)

var logger = zap.NewLogger()

func SendFriendRequest(c *gin.Context) {
    data, err := services.GetFriendModel(c)
    if err != nil {
        c.String(http.StatusBadRequest, err.Error())
        logger.Error(err.Error())
        return
    }

    if err := services.InsertIntoFriends(data); err != nil {
        c.String(http.StatusInternalServerError, err.Error())
        logger.Error(err.Error())
        return
    }

    logger.Info("Successfully sent friend request")
}

func AcceptFriendRequest(c *gin.Context) {
    data, err := services.GetFriendModel(c)
    if err != nil {
        c.String(http.StatusBadRequest, err.Error())
        logger.Error(err.Error())
        return
    }

    if err := services.AcceptFriendRequest(data); err != nil {
        c.String(http.StatusInternalServerError, err.Error())
        logger.Error(err.Error())
        return
    }

    logger.Info("Successfully accepted the friend request")
}

func RejectFriendRequest(c *gin.Context) {
    data, err := services.GetFriendModel(c)
    if err != nil {
        c.String(http.StatusBadRequest, err.Error())
        logger.Error(err.Error())
        return
    }

    if err := services.RejectFriendRequest(data); err != nil {
        c.String(http.StatusInternalServerError, err.Error())
        logger.Error(err.Error())
        return
    }

    logger.Info("Successfully rejected the friend request")
}

func RemoveFriend(c *gin.Context) {
    senderIDRaw, exists := c.Get("userID")
    if !exists {
        c.String(http.StatusBadRequest, "Cannot find user id")
        logger.Error("Cannot find user id")
        return
    }

    senderID, ok := senderIDRaw.(uint)
    if !ok {
        c.String(http.StatusBadRequest, "Invalid user id")
        logger.Error("Invalid user id")
        return
    }

    userToRemoveIDRaw := c.Param("friend_id")
    userToRemoveIDFoo, err := strconv.ParseUint(userToRemoveIDRaw, 10, 64)
    userToRemoveID := uint(userToRemoveIDFoo)
    if err != nil {
        c.String(http.StatusInternalServerError, err.Error())
        logger.Error(err.Error())
        return
    }

    if err := services.RemoveFriendFromFriends(senderID, userToRemoveID); err != nil {
        c.String(http.StatusInternalServerError, err.Error())
        logger.Error(err.Error())
        return
    }

    logger.Info("Successfully remove friend from friend list")
}

func GetFriends(c *gin.Context) {
    senderIDRaw, exists := c.Get("userID")
    if !exists {
        c.String(http.StatusBadRequest, "Cannot find user id")
        logger.Error("Cannot find user id")
        return
    }

    senderID, ok := senderIDRaw.(uint)
    if !ok {
        c.String(http.StatusBadRequest, "Invalid user id")
        logger.Error("Invalid user id")
        return
    }

    friends, err := services.GetFriends(senderID)
    if err != nil {
        c.String(http.StatusBadRequest, err.Error())
        logger.Error(err.Error())
        return
    }

    c.JSON(http.StatusFound, friends)
    logger.Info("Successfully get friends from friend list")
}

func GetFriendRequest(c *gin.Context) {
    receiverIDRaw, exists := c.Get("userID")
    if !exists {
        c.String(http.StatusBadRequest, "Cannot find user id")
        logger.Error("Cannot find user id")
        return
    }

    receiverID, ok := receiverIDRaw.(uint)
    if !ok {
        c.String(http.StatusBadRequest, "Invalid user id")
        logger.Error("Invalid user id")
        return
    }

    pendingReq, err := services.PendingRequests(receiverID)
    if err != nil {
        c.String(http.StatusBadRequest, err.Error())
        logger.Error(err.Error())
        return
    }

    c.JSON(http.StatusFound, pendingReq)
    logger.Info("Successfully get all pending friend requests")
}
