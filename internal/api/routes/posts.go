package routes

import (
	postsHandlers "github.com/jimtrung/go-nexus/internal/api/handlers/posts"
	"github.com/jimtrung/go-nexus/internal/middleware"
)

func SetUpPostsRoutes(r *Routes) {
    posts := r.Router.Group("/posts", middleware.RequireAuth)
    {
        posts.POST("/", postsHandlers.CreatePost)
        posts.GET("/")
        posts.GET("/:id")
        posts.DELETE("/:id")
        posts.PATCH("/:id")
        posts.POST("/:id/like")
        posts.GET("/:id/likes")
    }
}
