package v1

import (
	"net/http"
	"social_network_for_programmers/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	messenger := router.Group("/messages")
	{
		messenger.GET("/", h.services.Messenger.GetChatsHandler)
		messenger.GET("/page", func(c *gin.Context) {
			c.HTML(http.StatusOK, "test.html", gin.H{})
		})

		messages := messenger.Group("/ws")
		{
			messages.GET("/", h.services.Messenger.GetChatHandler)
			messages.GET("/:ChatId", h.services.Messenger.SendMessageHandler)
		}
	}

	authentification := router.Group("/auth")
	{
		authentification.GET("/", nil)
		authentification.POST("/", h.services.Authentication.CreateUser)
	}

	return router
}
