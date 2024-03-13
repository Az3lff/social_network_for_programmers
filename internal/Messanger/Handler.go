package Messanger

import (
	"social_network_for_programmers/internal/Handler"
	"github.com/gin-gonic/gin"
)

type handler struct {
}

func NewMessangerHandler() Handler.Handler {
	return &handler{}
}

func (h *handler) Register(router *gin.Engine) {
	MessangerApi := router.Group("/messages")
	MessangerApi.GET("/", h.GetChatsHandler)
	Messages := MessangerApi.Group("/:UserId")
	Messages.POST("/:message", h.SendMessageHandler)
	Messages.GET("/", h.GetChatHandler)
	// Messages.
}
