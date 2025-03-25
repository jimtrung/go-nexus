package routes

import (
	"github.com/jimtrung/go-nexus/internal/api/handlers/friend"
	"github.com/jimtrung/go-nexus/internal/infra/logger/zap"
	"github.com/jimtrung/go-nexus/internal/middleware"
	"github.com/jimtrung/go-nexus/internal/repository"
	"github.com/jimtrung/go-nexus/internal/services"
)

func (r *Routes) SetupFriendRoutes(logger *zap.Logger) {
	friendRepo := repository.NewFriendRepository(r.Conn)
	friendService := services.NewFriendService(friendRepo)
	friendHandler := friend.NewFriendHandler(friendService, logger)

	friendRouter := r.Router.Group("/friends", middleware.RequireAuth)
	{
		friendRouter.GET("/", friendHandler.GetAllFriends)
		friendRouter.GET("/requests", friendHandler.GetPendingRequests)
		friendRouter.POST("/request", friendHandler.CreateRequest)
		friendRouter.POST("/accept", friendHandler.AcceptRequest)
		friendRouter.POST("/reject", friendHandler.RejectRequest)
		friendRouter.DELETE("/cancel/:receiver_id", friendHandler.CancelRequest)
		friendRouter.DELETE("/remove/:friend_id", friendHandler.RemoveFriend)
	}
}
