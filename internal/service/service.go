package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"social_network_for_programmers/internal/entity/users"
	"social_network_for_programmers/internal/repository"
	"social_network_for_programmers/pkg/auth/tokenutil"
)

type Users interface {
	SignUp(c context.Context, user *users.UserSignUp) error
	SignIn(ctx context.Context, user *users.UserSignIn) (string, error)
}

type Messenger interface {
	GetConnChatHandler(c *gin.Context)
	GetChatHandler(c *gin.Context)
	CreateChatHandler(c *gin.Context)
	GetAllChatsHandler(c *gin.Context)
}

type Services struct {
	Users     Users
	Messenger Messenger
}

func NewServices(repos *repository.Repositories, tokenManager tokenutil.TokenManager) *Services {
	usersService := NewUsersService(repos.Users, tokenManager)
	messengerService := NewMessengerService(repos.Messenger)

	return &Services{
		usersService,
		messengerService,
	}
}
