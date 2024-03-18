package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"social_network_for_programmers/internal/entity"
)

type Authentication interface {
	CreateUser(user *entity.UsersSignUpInput) error
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

func NewRepositories(db *pgx.Conn) *Repositories {
	return &Repositories{
		NewAuthenticationRepo(db),
		NewMessengerRepo(),
	}
}
