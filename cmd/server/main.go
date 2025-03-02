package main

import (
	"github.com/jimtrung/go-nexus/internal/api"
	authHandlers "github.com/jimtrung/go-nexus/internal/api/handlers/auth"
	"github.com/jimtrung/go-nexus/internal/infra/db"
	"github.com/jimtrung/go-nexus/internal/infra/env"
	"github.com/jimtrung/go-nexus/internal/infra/logger/zap"
)

func main() {
	logger := zap.NewLogger()
	if err := env.SetupEnv(); err != nil {
		logger.Error("Error setting up environment", err)
		return
	}

	database := db.ConnectToDatabase()
    authHandlers.NewOAuth()
	server := api.NewServer("debug")

	port, err := env.GetPort()
	if err != nil {
		logger.Error("Error getting port from .env", err)
		return
	}

	server.StartServer(database, port)
}
