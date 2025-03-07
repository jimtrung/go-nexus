package routes

import (
	friendsHandler "github.com/jimtrung/go-nexus/internal/api/handlers/friends"
	"github.com/jimtrung/go-nexus/internal/middleware"
)

func SetupFriendsRoutes(r *Routes) {
	friends := r.Router.Group("/friends")
	{
        friends.GET("/")
        friends.GET("/requests")
        friends.POST("/request", middleware.RequireAuth, friendsHandler.SendFriendRequest)
        friends.POST("/accept")
        friends.POST("/reject")
        friends.DELETE("/remove/:friend_id")
	}
}
