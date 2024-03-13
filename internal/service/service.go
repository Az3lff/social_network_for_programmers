package service

import (
	"github.com/gin-gonic/gin"
	"social_network_for_programmers/internal/repository"
)

type Authentication interface {
	CreateUser(c *gin.Context)
}

type Messenger interface {
	GetChatsHandler(c *gin.Context)
	SendMessageHandler(c *gin.Context)
	GetChatHandler(c *gin.Context)
}

type Services struct {
	Authentication Authentication
	Messenger      Messenger
}

func NewServices(repos *repository.Repositories) *Services {
	authenticationService := NewAuthenticationService(repos.Authentication)
	messengerService := NewMessengerService(repos.Messenger)

	return &Services{
		authenticationService,
		messengerService,
	}
}
