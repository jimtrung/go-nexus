package models

import "time"

type Role string

const (
	Admin     Role = "admin"
	Moderator Role = "moderator"
	Client    Role = "client"
)

type User struct {
	UserID    uint      `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Username  string    `gorm:"type:varchar(30);not null;unique" json:"username"`
	Email     string    `gorm:"type:varchar(100);not null;unique" json:"email"`
	Password  string    `gorm:"type:varchar(100);not null" json:"-"`
	Role      Role      `gorm:"type:varchar(20);default:'client'" json:"role"`
	Verified  bool      `gorm:"default:false" json:"verified"`
	CreatedAt time.Time `gorm:"type:timestamptz;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamptz;autoUpdateTime" json:"updated_at"`
}
