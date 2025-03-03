package services

import (
	"github.com/jimtrung/go-nexus/internal/domain/models"
)

func SignupIfNotExist(email string) (models.User, error) {
	userInfo, err := GetUserByEmail(email)

	if err == nil {
		return userInfo, nil
	}

	randomPassword := GenerateRandomPassword()
	hashedPassword, err := HashPassword(randomPassword)
	if err != nil {
		return models.User{}, err
	}

	userInfo = models.User{
		Username: email,
		Email:    email,
		Verified: true,
		Password: hashedPassword,
	}
    if err := InsertIntoUsers(userInfo); err != nil {
		return models.User{}, err
    }

	return userInfo, nil
}
