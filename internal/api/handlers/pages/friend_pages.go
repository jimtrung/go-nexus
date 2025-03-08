package pageshandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/go-nexus/internal/api/handlers"
	"github.com/jimtrung/go-nexus/internal/services"
	friendComponent "github.com/jimtrung/go-nexus/templates/components/friend"
)

func RenderShowPage(c *gin.Context) {
    userIDRaw, exists := c.Get("userID")
    if !exists {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Cannot find user id",
        })
        logger.Error("Cannot find user id")
        return
    }

    userID, ok := userIDRaw.(uint)
    if !ok {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": "Invalid user id",
        })
        logger.Error("Invalid user id")
        return
    }

    friends, err := services.GetFriends(userID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        logger.Error(err.Error())
        return
    }

    requests, err := services.PendingRequests(userID)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error": err.Error(),
        })
        logger.Error(err.Error())
        return
    }

    if err := handlers.Render(c, friendComponent.Show(friends, requests)); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "Failed to render the page",
        })
        logger.Error("Failed to render the page", err.Error())
        return
    }
    logger.Info("Show page rendered successfully")
}
