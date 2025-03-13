package routes

import (
	friendsHandler "github.com/jimtrung/go-nexus/internal/api/handlers/friends"
	"github.com/jimtrung/go-nexus/internal/middleware"
)

func SetupFriendsRoutes(r *Routes) {
	friends := r.Router.Group("/friends", middleware.RequireAuth)
	{
        friends.GET("/", friendsHandler.GetFriends)
        friends.GET("/requests", friendsHandler.GetFriendRequest)
        friends.POST("/request", friendsHandler.SendFriendRequest)
        friends.POST("/accept", friendsHandler.AcceptFriendRequest)
        friends.POST("/reject", friendsHandler.RejectFriendRequest)
        friends.DELETE("/remove/:friend_id", friendsHandler.RemoveFriend)
	}
}
