package models

import (
	"time"
)

type Friend struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	SenderID   uint      `gorm:"not null" json:"sender_id"`
	ReceiverID uint      `gorm:"not null" json:"receiver_id"`
	Status     Status    `gorm:"type:varchar(20);default:'pending'" json:"status"`
	CreatedAt  time.Time `gorm:"type:timestamptz;autoCreateTime" json:"created_at"`
	UpdatedAt  time.Time `gorm:"type:timestamptz;autoCreateTime" json:"updated_at"`
}

type FriendData struct {
	UserID   uint
	Username string
}
