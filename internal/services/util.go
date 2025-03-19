package services

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func GenerateRandomPassword() string {
	legitChars := "qwertyuiopasdfghjklzxcvbnm1234567890@."

	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	res := make([]byte, 8)
	for i := range res {
		res[i] = legitChars[seededRand.Intn(len(legitChars))]
	}
	return string(res)
}

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", fmt.Errorf("Failed to hash password")
	}
	return string(hashed), nil
}
