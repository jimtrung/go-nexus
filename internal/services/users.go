package services

import (
	"fmt"

	"github.com/jimtrung/go-nexus/internal/domain/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func InsertIntoUsers(conn *gorm.DB, user models.User) error {
    result := conn.Create(&user)

    return result.Error
}

func HashPassword(password string) (string, error) {
    hashed, err := bcrypt.GenerateFromPassword([]byte(password), 8)
    if err != nil {
        return "", fmt.Errorf("Failed to hash password")
    }

    return string(hashed), nil
}
