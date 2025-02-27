package db

import (
	"log"

	"github.com/jimtrung/go-nexus/internal/infra/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDatabase() *gorm.DB {
	dbUrl, err := env.GetDBURL()
	if err != nil {
		log.Fatalf("Error getting database url: %s", err)
	}

	conn, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %s", err)
	}
    DB = conn

	return conn
}
