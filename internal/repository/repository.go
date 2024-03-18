package repository

import "github.com/gin-gonic/gin"

type Authentication interface {
	CreateUser(c *gin.Context)
}

type Messenger interface {
	GetChatsHandler()
	SendMessageHandler()
	GetChatHandler()
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
