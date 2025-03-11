package models

import "time"

type Post struct {
    PostID    uint      `gorm:"primaryKey" json:"post_id"`
    UserID    uint      `gorm:"index" json:"user_id"`
    Content   string    `gorm:"type:text" json:"content"`
    Likes     int       `gorm:"default:0" json:"likes"`
	CreatedAt time.Time `gorm:"type:timestamptz;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamptz;autoUpdateTime" json:"updated_at"`
    Comments  []Comment `gorm:"foreignKey:PostID" json:"comments"`
}
