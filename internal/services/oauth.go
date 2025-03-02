package services

import (
	"errors"

	"github.com/jimtrung/go-nexus/internal/domain/models"
	"gorm.io/gorm"
)

func SignupIfNotExist(email string) (models.User, error) {
    userInfo, err := GetUserByEmail(email)
    if err == nil {
        return userInfo, nil
    }

    if !errors.Is(err, gorm.ErrRecordNotFound) {
        return models.User{}, nil
    }

	randomPassword := GenerateRandomPassword()
    hashedPassword, err := HashPassword(randomPassword)
    if err != nil {
        return models.User{}, err
    }

    userInfo = models.User{
        Username: email,
        Email: email,
        Password: hashedPassword,
    }
    err = InsertIntoUsers(userInfo)

    return userInfo, err
}
