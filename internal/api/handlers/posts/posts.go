package postshandlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/go-nexus/internal/domain/models"
	"github.com/jimtrung/go-nexus/internal/infra/logger/zap"
	"github.com/jimtrung/go-nexus/internal/services"
)

var logger = zap.NewLogger();

func CreatePost(c *gin.Context) {
    userID, err := services.GetUserID(c);
    if err != nil {
        c.String(http.StatusBadRequest, err.Error())
        logger.Error(err.Error())
        return
    }

    var createReq models.Post
    if err := c.Bind(&createReq); err != nil {
        c.String(http.StatusBadRequest, "Wrong JSON format")
        logger.Error("Wrong JSON format")
        return
    }
    createReq.UserID = userID

    if err := services.InsertIntoPosts(createReq); err != nil {
        c.String(http.StatusInternalServerError, err.Error())
        logger.Error(err.Error())
        return
    }

    logger.Info("Successfully added post to database")
}
