package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"social_network_for_programmers/internal/config"
	"social_network_for_programmers/internal/entity/auth_entity"
	"social_network_for_programmers/internal/repository"
	"social_network_for_programmers/pkg/auth"
)

type Auth interface {
	SignUp(ctx context.Context, user *auth_entity.UserSignUp) error
	SignIn(ctx context.Context, user *auth_entity.UserSignIn) (string, error)
	RestoreAccount(ctx context.Context, email string, cfg *config.AuthEmail) error
	CheckRestoreCode(ctx context.Context, req *auth_entity.RestoreAccessRequest) error
	UpdatePassword(ctx context.Context, user *auth_entity.UserUpdatePassword, ath *config.AuthEmail) error
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

func NewServices(repos *repository.Repositories, tokenManager auth.TokenManager, cache *redis.Client) *Services {
	authService := NewAuthService(repos.Auth, tokenManager, cache)
	messengerService := NewMessengerService(repos.Messenger)

	return &Services{
		authService,
		messengerService,
	}
}
