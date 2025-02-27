package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type Routes struct {
	Router   *gin.Engine
	Database *pgx.Conn
}

func (r *Routes) SetupRoutes() {
}

