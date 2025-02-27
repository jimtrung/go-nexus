package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jimtrung/go-nexus/internal/infra/env"
)

func ConnectToDatabase() *pgx.Conn {
    dbUrl, err := env.GetDBURL()
    if err != nil {
        log.Fatalf("Error getting database url: %s", err)
    }

    conn, err := pgx.Connect(context.Background(), dbUrl)
    if err != nil {
        log.Fatalf("Error connecting to database: %s", err)
    }

    return conn
}

func CloseConnection(conn *pgx.Conn) {
    conn.Close(context.Background())
}

