package v1

import (
	"github.com/gin-gonic/gin"
	"social_network_for_programmers/internal/service"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	messenger := router.Group("/messages")
	{
		messenger.GET("/", h.services.Messenger.GetChatsHandler)

		messages := messenger.Group("/:UserId")
		{
			messages.GET("/", h.services.Messenger.GetChatHandler)
			messages.POST("/:message", h.services.Messenger.SendMessageHandler)
		}
	}

	authentification := router.Group("/auth")
	{
		authentification.GET("/", nil)
		authentification.POST("/", h.services.Authentication.CreateUser)
	}

	return router
}
