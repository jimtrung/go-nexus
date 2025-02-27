package services

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateSignedToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 4).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	return tokenString, err
}
