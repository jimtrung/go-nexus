package services

import (
	"fmt"

	"github.com/jimtrung/go-nexus/internal/domain/models"
	"github.com/jimtrung/go-nexus/internal/infra/db"
	"golang.org/x/crypto/bcrypt"
)

func InsertIntoUsers(user models.User) error {
    result := db.DB.Create(&user)
    if result.Error != nil {
        return fmt.Errorf("Username/Email is already used")
    }

    return nil
}

func GetUserByUsername(username string) (models.User, error) {
    var res models.User

    result := db.DB.Select(
        "username", "email", "created_at",
    ).Where("username = ?", username).Find(&res)
    if result.Error != nil {
        return res, fmt.Errorf("Cannot find user with email %s", username)
    }

    return res, nil
}

func GetUserByEmail(email string) (models.User, error) {
    var res models.User

    result := db.DB.Select(
        "username", "email", "created_at",
    ).Where("email = ?", email).Find(&res)
    if result.RowsAffected == 0 {
        return res, fmt.Errorf("Cannot find user with email %s", email)
    }

    return res, nil
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
        return fmt.Errorf("Wrong password")
    }

    return nil
}

func AddTokenToUser(email, token string) error {
    result := db.DB.Table("users").Where("email = ?", email).Update("token", token)

    return result.Error
}

func RemoveToken(token string) {
    db.DB.Table("users").Where("token = ?", token).Update("token", "")
}

func VerifyUser(token string) error {
    if token == "" {
        return fmt.Errorf("Cannot find the token")
    }

    result := db.DB.Table("users").Where("token = ?", token).Update("verified", "true").Update("token", "")

    return result.Error
}

func ResetPassword(token, newPassword string) error {
    if token == "" {
        return fmt.Errorf("Cannot find the token")
    }

    hashedPassword, err := HashPassword(newPassword)
    if err != nil {
        return err
    }

    result := db.DB.Table("users").Where("token = ?", token).Update("password", hashedPassword).Update("token", "")

    return result.Error
}
