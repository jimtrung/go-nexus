package services

import (
	"fmt"

	"github.com/jimtrung/go-nexus/internal/domain/models"
	"github.com/jimtrung/go-nexus/internal/infra/db"
)

func InsertIntoPosts(createReq models.Post) error {
    result := db.DB.Create(&createReq)
    if result.RowsAffected == 0 {
        return fmt.Errorf("Failed to insert into database")
    }
    return nil
}
