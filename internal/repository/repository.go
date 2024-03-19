package repository

import (
	//"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"social_network_for_programmers/internal/entity"
)

type Authentication interface {
	CreateUser(user *entity.UsersSignUpInput) error
	CheckUser(login, password string) error
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

func NewRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		NewAuthenticationRepo(db),
		NewMessengerRepo(),
	}
}
