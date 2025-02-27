package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jimtrung/go-nexus/templates/layout"
)

func main() {
    r := gin.Default()

    r.GET("/greet", func(ctx *gin.Context) {
        ctx.JSON(http.StatusOK, gin.H{
            "message": "Hello, World",
        })
    })

    r.GET("/templ", func(ctx *gin.Context) {
        layout.Base().Render(ctx, ctx.Writer)
    })

    fmt.Println("Server is running on port 8080")
    r.Run("127.0.0.1:8080")
}
