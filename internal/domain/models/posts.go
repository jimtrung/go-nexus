package models

import "time"

type Post struct {
	ID        uint      `gorm:"primaryKey"`
	UserID    uint      `gorm:"index"`
	Content   string    `gorm:"type:text"`
	Likes     int       `gorm:"default:0"`
	CreatedAt time.Time `gorm:"type:timestamptz;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamptz;autoUpdateTime" json:"updated_at"`
	Comments  []Comment `gorm:"foreignKey:PostID"`
}
