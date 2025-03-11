package routes

import postsHandlers "github.com/jimtrung/go-nexus/internal/api/handlers/posts"

func SetUpPostsRoutes(r *Routes) {
    posts := r.Router.Group("/posts")
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
