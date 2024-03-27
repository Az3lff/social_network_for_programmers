package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	messengerModels "social_network_for_programmers/internal/entity/messenger"
	"social_network_for_programmers/internal/entity/users"
)

type Auth interface {
	Create(ctx context.Context, user *users.UserSignUp) error
	GetByEmail(ctx context.Context, email string) (user *users.UserRepo, err error)
	FindByEmail(ctx context.Context, email string) error
}

type Messenger interface {
	SaveMessage(ChatId string, mess *messengerModels.Message) error
	GetMessages(ChatId string, messages *[]messengerModels.Message) error
	CreateChat(senderId string, recipientId string) (string, error)
	GetAllChats(UserId string) ([]string, error)
}

type Repositories struct {
	Auth      Auth
	Messenger Messenger
}

func NewRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		NewAuthRepo(db),
		NewMessengerRepo(db),
	}
}
