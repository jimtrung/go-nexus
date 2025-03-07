package services

import (
	"fmt"

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
    //  UPDATE friends
    //  SET status = 'accepted'
    //  WHERE (sender_id = 5 AND receiver_id = 6)
    //     OR (sender_id = 6 AND receiver_id = 5);
    return nil
}
