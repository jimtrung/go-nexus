package services

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/go-nexus/internal/domain/models"
	"github.com/jimtrung/go-nexus/internal/infra/db"
)

func InsertIntoFriends(friendReq models.Friend) error {
    result := db.DB.Create(&friendReq)
    if result.Error != nil {
        return fmt.Errorf("Request already existed")
    }

    return nil
}

func AcceptFriendRequest(req models.Friend) error {
    result := db.DB.Table("friends").Where(
        "(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)",
        req.SenderID, req.ReceiverID, req.ReceiverID, req.SenderID,
    ).Update("status", "accepted")
    if result.RowsAffected == 0 {
        return fmt.Errorf("Failed to accept the friend request")
    }
    return nil
}

func RejectFriendRequest(req models.Friend) error {
    result := db.DB.Table("friends").Where(
        "(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)",
        req.SenderID, req.ReceiverID, req.ReceiverID, req.SenderID,
    ).Update("status", "rejected")
    if result.RowsAffected == 0 {
        return fmt.Errorf("Failed to reject the friend request")
    }
    return nil
}

func GetFriendModel(c *gin.Context) (models.Friend, error) {
    senderIdRaw, exists := c.Get("userID")
    if !exists {
        return models.Friend{}, fmt.Errorf("User ID not existed")
    }

    senderID, ok := senderIdRaw.(uint)
    if !ok {
        return models.Friend{}, fmt.Errorf("Not a valid user ID")
    }

    var req models.Friend
    if err := c.Bind(&req); err != nil {
        return models.Friend{}, fmt.Errorf("Wrong JSON format")
    }

    data := models.Friend {
        SenderID: senderID,
        ReceiverID: req.ReceiverID,
    }

    return data, nil
}

func RemoveFriendFromFriends(senderID, receiverID uint) error {
    result := db.DB.Where(
`((sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?))
AND status = ?`,
        senderID, receiverID, receiverID, senderID, "accepted",
    ).Delete(&models.Friend{})
    if result.RowsAffected == 0 {
        return fmt.Errorf("Failed to delete friend")
    }

    return nil
}

func GetFriends(userID uint) ([]models.Friend, error) {
    var friends []models.Friend
    result := db.DB.Where(
        "(sender_id = ? OR receiver_id = ?) AND status = ?",
        userID, userID, "accepted",
    ).Find(&friends)
    if result.Error != nil {
        return []models.Friend{}, result.Error
    }

    return friends, nil
}

func PendingRequests(userID uint) ([]models.Friend, error) {
    var friends []models.Friend
    result := db.DB.Where(
        "receiver_id = ? AND status = ?",
        userID, "pending",
    ).Find(&friends)
    if result.Error != nil {
        return []models.Friend{}, result.Error
    }

    return friends, nil
}
