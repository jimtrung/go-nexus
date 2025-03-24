package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jimtrung/go-nexus/internal/infra/logger/zap"
)

type Routes struct {
	Router      *gin.Engine
	Conn        *pgx.Conn
}

func (r *Routes) SetupRoutes(logger *zap.Logger) {
	r.SetupAuthRoutes(logger)
    r.SetupPageRoutes(logger)
	r.SetupFriendRoutes(logger)
}
