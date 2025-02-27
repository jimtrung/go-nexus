package env

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func SetupEnv() error {
    if err := godotenv.Load(); err != nil {
        return fmt.Errorf("Failed to load .env file")
    }

    return nil
}

func GetPort() (string, error) {
    port := os.Getenv("PORT")
    if port == "" {
        return port, fmt.Errorf("Failed to get PORT from .env")
    }

    return port, nil
}

func GetDBURL() (string, error) {
    db := os.Getenv("DB_URL")
    if db == "" {
        return db, fmt.Errorf("Failed to get DB URL from .env")
    }

    return db, nil
}

