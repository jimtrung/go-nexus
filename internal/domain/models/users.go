package models

import "time"

type User struct {
	UserID    uint      `gorm:"primaryKey;autoIncrement"`
	Username  string    `gorm:"type:varchar(30);not null;unique"`
	Email     string    `gorm:"type:varchar(100);not null;unique"`
	Password  string    `gorm:"type:varchar(100);not null"`
	CreatedAt time.Time `gorm:"type:timestamptz;autoCreateTime"`
	UpdatedAt time.Time `gorm:"type:timestamptz;autoUpdateTime"`
}
