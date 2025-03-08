package routes

import (
	friendsHandler "github.com/jimtrung/go-nexus/internal/api/handlers/friends"
	"github.com/jimtrung/go-nexus/internal/middleware"
)

func SetupFriendsRoutes(r *Routes) {
	friends := r.Router.Group("/friends")
	{
        friends.GET("/", middleware.RequireAuth, friendsHandler.GetFriends)
        friends.GET("/requests", middleware.RequireAuth, friendsHandler.GetFriendRequest)
        friends.POST("/request", middleware.RequireAuth, friendsHandler.SendFriendRequest)
        friends.POST("/accept", middleware.RequireAuth, friendsHandler.AcceptFriendRequest)
        friends.POST("/reject", middleware.RequireAuth, friendsHandler.RejectFriendRequest)
        friends.DELETE("/remove/:friend_id", middleware.RequireAuth, friendsHandler.RemoveFriend)
	}
}
