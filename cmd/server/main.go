package main

import (
	"os"

	"github.com/jimtrung/go-nexus/internal/api"
	"github.com/jimtrung/go-nexus/internal/infra/db"
	"github.com/jimtrung/go-nexus/internal/infra/env"
	"github.com/jimtrung/go-nexus/internal/infra/logger/zap"
	"github.com/jimtrung/go-nexus/internal/middleware"
)

func main() {
	logger := zap.NewLogger()
	if err := env.SetupEnv(); err != nil {
		logger.Error("Error setting up environment", err)
		return
	}

	conn := db.ConnectToDatabase()
    middleware.NewOAuth()
	server := api.NewServer("debug")

    server.StartServer(conn, os.Getenv("PORT"), logger)
}
