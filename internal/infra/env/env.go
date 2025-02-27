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
	db := ""

	host := os.Getenv("DB_HOST")
	if host == "" {
		return db, fmt.Errorf("Failed to get database host from .env")
	}

	username := os.Getenv("DB_USERNAME")
	if username == "" {
		return db, fmt.Errorf("Failed to get database username from .env")
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		return db, fmt.Errorf("Failed to get database password from .env")
	}

	name := os.Getenv("DB_NAME")
	if name == "" {
		return db, fmt.Errorf("Failed to get database name from .env")
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		return db, fmt.Errorf("Failed to get database port from .env")
	}

    db = fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
        host, username, password, name, port,
    )

	return db, nil
}
