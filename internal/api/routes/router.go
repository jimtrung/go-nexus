package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Routes struct {
	Router   *gin.Engine
	Database *gorm.DB
}

func (r *Routes) SetupRoutes() {
    SetupAuthRoutes(r)
}
