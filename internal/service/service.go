package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"social_network_for_programmers/internal/config"
	"social_network_for_programmers/internal/entity/users"
	"social_network_for_programmers/internal/repository"
	"social_network_for_programmers/pkg/auth"
)

type Auth interface {
	SignUp(ctx context.Context, user *users.UserSignUp) error
	SignIn(ctx context.Context, user *users.UserSignIn) (string, error)
	RestoreAccount(ctx context.Context, email string, cfg *config.AuthEmail) error
}

type Messenger interface {
	GetConnChatHandler(c *gin.Context)
	GetChatHandler(c *gin.Context)
	CreateChatHandler(c *gin.Context)
	GetAllChatsHandler(c *gin.Context)
}

type Services struct {
	Auth      Auth
	Messenger Messenger
}

func NewServices(repos *repository.Repositories, tokenManager auth.TokenManager) *Services {
	authService := NewAuthService(repos.Auth, tokenManager)
	messengerService := NewMessengerService(repos.Messenger)

	return &Services{
		authService,
		messengerService,
	}
}
