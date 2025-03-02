package services

import (
	"fmt"

	"github.com/jimtrung/go-nexus/internal/domain/models"
	"github.com/jimtrung/go-nexus/internal/infra/db"
	"golang.org/x/crypto/bcrypt"
)

func InsertIntoUsers(user models.User) error {
    result := db.DB.Create(&user)

    return result.Error
}

func GetUserByUsername(username string) (models.User, error) {
    var res models.User

    result := db.DB.Select(
        "username", "email", "created_at",
    ).Where("username = ?", username).Find(&res)

    return res, result.Error
}

func HashPassword(password string) (string, error) {
    hashed, err := bcrypt.GenerateFromPassword([]byte(password), 8)
    if err != nil {
        return "", fmt.Errorf("Failed to hash password")
    }

    return string(hashed), nil
}

func IsValidUser(user models.User) error {
    var res models.User

    result := db.DB.Select("password").Where("username = ?", user.Username).Find(&res)
    if result.Error != nil {
        return result.Error
    }

    if err := bcrypt.CompareHashAndPassword(
        []byte(res.Password), []byte(user.Password),
    ); err != nil {
        return err
    }

    return nil
}
