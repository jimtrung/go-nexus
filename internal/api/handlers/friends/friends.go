package friends

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/go-nexus/internal/domain/models"
	"github.com/jimtrung/go-nexus/internal/infra/logger/zap"
	"github.com/jimtrung/go-nexus/internal/services"
)

var logger = zap.NewLogger()

func SendFriendRequest(c *gin.Context) {
    senderIdRaw, exists := c.Get("userID")
    if !exists {
        c.String(http.StatusBadRequest, "Cannot find user ID")
        logger.Error("Cannot find user ID")
        return
    }

    senderID, ok := senderIdRaw.(uint)
    if !ok {
        c.String(http.StatusBadRequest, "Not a valid user ID")
        logger.Error("Not a valid user ID")
        return
    }

    var receiver models.Friend
    if err := c.Bind(&receiver); err != nil {
        c.String(http.StatusBadRequest, "Wrong JSON format")
        logger.Error("Wrong JSON format")
        return
    }

    friendReq := models.Friend {
        SenderID: senderID,
        ReceiverID: receiver.ReceiverID,
        Status: models.Pending,
    }

    if err := services.InsertIntoFriends(friendReq); err != nil {
        c.String(http.StatusInternalServerError, err.Error())
        logger.Error(err.Error())
        return
    }

    logger.Info("Successfully sent friend request")
}

func AcceptFriendRequest(c *gin.Context) {
    var other models.Friend

    if err := c.Bind(&other); err != nil {
        c.String(http.StatusBadRequest, "Wrong JSON format")
        logger.Error("Wrong JSON format")
        return
    }

    // Accept the friend request

    logger.Info("Successfully accepted the friend request")
}
