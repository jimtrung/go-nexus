package pageshandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/go-nexus/internal/api/handlers"
	friendComponent "github.com/jimtrung/go-nexus/templates/components/friend"
)

func RenderShowPage(c *gin.Context) {
    if err := handlers.Render(c, friendComponent.Show()); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to render the page",
        })
        logger.Error("Failed to render the page", err.Error())
        return
    }
    logger.Info("Show page rendered successfully")
}
