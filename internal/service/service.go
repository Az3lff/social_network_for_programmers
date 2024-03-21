package service

import (
	"github.com/gin-gonic/gin"
	"social_network_for_programmers/internal/repository"
	"social_network_for_programmers/pkg/auth"
)

type Users interface {
	SignUp(c *gin.Context)
	SignUpPage(c *gin.Context)
	SignInPage(c *gin.Context)
	SignIn(c *gin.Context)
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

func NewServices(repos *repository.Repositories, tokenManager auth.TokenManager) *Services {
	usersService := NewUsersService(repos.Users, tokenManager)
	messengerService := NewMessengerService(repos.Messenger)

	return &Services{
		usersService,
		messengerService,
	}
}
