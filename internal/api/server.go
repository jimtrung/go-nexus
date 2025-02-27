package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/jimtrung/go-nexus/internal/api/routes"
)

type Server struct {
	router *gin.Engine
}

func NewServer(mode string) *Server {
    setServerMode(mode)
    server := &Server{}
    return server
}

func (server *Server) setupRouter(database *pgx.Conn) {
	router := gin.Default()
    router.SetTrustedProxies([]string{"127.0.0.1"})
    routes := routes.Routes{Router: router, Database: database}
    routes.SetupRoutes()
    server.router = router
}

func (server *Server) StartServer(database *pgx.Conn, port string) error {
    address := "127.0.0.1:" + port
	server.setupRouter(database)
	return server.router.Run(address)
}

func setServerMode(mode string) {
    if mode == gin.ReleaseMode {
        gin.SetMode(gin.ReleaseMode)
    } else if mode == gin.DebugMode {
        gin.SetMode(gin.DebugMode)
    }
}

