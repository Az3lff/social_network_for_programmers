package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	messengerModels "social_network_for_programmers/internal/entity/messenger"
	"social_network_for_programmers/internal/entity/users"
)

type Users interface {
	Create(ctx context.Context, user *users.UserSignUp) error
	Find(ctx context.Context, user *users.UserSignIn) (id string, err error)
}

type Messenger interface {
	SaveMessage(ChatId string, mess *messengerModels.Message) error
	GetMessages(ChatId string, messages *[]messengerModels.Message) error
	CreateChat(senderId string, recipientId string) (string, error)
	GetAllChats(UserId string) ([]string, error)
}

type Repositories struct {
	Users     Users
	Messenger Messenger
}

func NewRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		NewUsersRepo(db),
		NewMessengerRepo(db),
	}
}
