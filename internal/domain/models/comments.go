package models

import "time"

type Comment struct {
	CommentID uint      `gorm:"primaryKey" json:"comment_id"`
	UserID    uint      `gorm:"primaryKey;autoIncrement" json:"user_id"`
	PostID    uint      `gorm:"index" json:"post_id"`
	Content   string    `gorm:"type:text" json:"content"`
	CreatedAt time.Time `gorm:"type:timestamptz;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamptz;autoUpdateTime" json:"updated_at"`
}
