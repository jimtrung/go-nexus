package env

import (
	"fmt"

	"github.com/joho/godotenv"
)

func SetupEnv() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("Failed to load .env file")
	}
	return nil
}
