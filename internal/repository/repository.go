package repository

import (
	"github.com/gin-gonic/gin"
	"social_network_for_programmers/internal/entity"
)

type Authentication interface {
	CreateUser(user *entity.User) error
	CheckUser(login, password string) error
}

type Messenger interface {
	GetChatsHandler(c *gin.Context)
	SendMessageHandler(c *gin.Context)
	GetChatHandler(c *gin.Context)
}

type Repositories struct {
	Authentication Authentication
	Messenger      Messenger
}

func NewRepositories() *Repositories {
	return &Repositories{
		NewAuthenticationRepo(),
		NewMessengerRepo(),
	}
}
