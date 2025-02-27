package routes

import "github.com/jimtrung/go-nexus/internal/api/handlers/pages"

func SetupPagesRoutes(r *Routes) {
    p := r.Router.Group("/p")
    {
        user := p.Group("/user")
        {
            user.GET("/signup", pages.RenderSignupPage)
        }
    }
}
