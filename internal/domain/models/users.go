package models

import "time"

type User struct {
	UserID    uint      `gorm:"primaryKey;autoIncrement" json:"user_id"`
	Username  string    `gorm:"type:varchar(30);not null;unique" json:"username"`
	Email     string    `gorm:"type:varchar(100);not null;unique" json:"email"`
	Password  string    `gorm:"type:varchar(100);not null" json:"password"`
	Role      Role      `gorm:"type:varchar(20);default:'user'" json:"role"`
	Token     string    `gorm:"type:varchar(50);default:''" json:"token"`
	Verified  bool      `gorm:"default:false" json:"verified"`
	CreatedAt time.Time `gorm:"type:timestamptz;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamptz;autoUpdateTime" json:"updated_at"`
}

type ResetPassword struct {
	Token           string `gorm:"type:varchar(50);default:''" json:"token"`
	Password        string `gorm:"type:varchar(100);not null" json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}
