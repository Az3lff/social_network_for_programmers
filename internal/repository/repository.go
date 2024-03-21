package repository

import (
	//"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"social_network_for_programmers/internal/entity"
	messengerModels "social_network_for_programmers/internal/entity/messenger"
)

type Authentication interface {
	CreateUser(user *entity.UsersSignUpInput) error
	CheckUser(login, password string) error
}

type Messenger interface {
	SaveMessage(ChatId string, mess *messengerModels.Message) error
	GetMessages(ChatId string, messages *[]messengerModels.Message) error
	CreateChat(senderId string, recipientId string) (string, error)
	GetAllChats(UserId string) ([]string, error)
}

type Repositories struct {
	Authentication Authentication
	Messenger      Messenger
}

func NewRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		NewAuthenticationRepo(db),
		NewMessengerRepo(db),
	}
}
