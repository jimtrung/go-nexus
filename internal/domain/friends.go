package domain

import "time"

type Friend struct {
	FriendID   uint    `json:"friend_id"`
	SenderID   uint    `json:"sender_id"`
	ReceiverID uint    `json:"receiver_id"`
	Status     Status  `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}