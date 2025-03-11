package services

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
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

func GetUserID(c *gin.Context) (uint, error) {
    rawUserID, exists := c.Get("userID")
    if !exists {
        return 0, fmt.Errorf("User ID not existed")
    }

    userID, ok := rawUserID.(uint)
    if !ok {
        return 0, fmt.Errorf("Not a valid user ID")
    }

    return userID, nil
}
