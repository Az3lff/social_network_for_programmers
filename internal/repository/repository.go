package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"social_network_for_programmers/internal/entity/auth_entity"
	messengerModels "social_network_for_programmers/internal/entity/messenger"
)

type Auth interface {
	Create(ctx context.Context, user *auth_entity.UserSignUp) error
	GetByEmail(ctx context.Context, email string) (user *auth_entity.UserRepo, err error)
	FindByEmail(ctx context.Context, email string) error
	UpdatePasswordByEmail(ctx context.Context, user *auth_entity.UserUpdatePassword) error
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
