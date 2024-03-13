package repository

import "github.com/gin-gonic/gin"

type Authentication interface {
	CreateUser(c *gin.Context)
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
