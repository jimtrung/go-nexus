package services

import (
	"math/rand"
	"time"
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
