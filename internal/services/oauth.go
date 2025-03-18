package services

import (
	"github.com/jimtrung/go-nexus/internal/domain"
)

func SignupIfNotExist(email string) (domain.User, error) {
	userInfo, err := GetUserByEmail(email)

	if err == nil {
		return userInfo, nil
	}

	randomPassword := GenerateRandomPassword()
	hashedPassword, err := HashPassword(randomPassword)
	if err != nil {
		return domain.User{}, err
	}

	userInfo = domain.User{
		Username: email,
		Email:    email,
		Verified: true,
		Password: hashedPassword,
	}
    if err := InsertIntoUsers(userInfo); err != nil {
		return domain.User{}, err
    }

	return userInfo, nil
}
