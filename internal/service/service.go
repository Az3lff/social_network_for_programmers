package service

import (
	"github.com/gin-gonic/gin"
	"social_network_for_programmers/internal/repository"
	"social_network_for_programmers/pkg/auth"
)

type Authentication interface {
	SignUp(c *gin.Context)
	SignUpPage(c *gin.Context)
	SignInPage(c *gin.Context)
	SignIn(c *gin.Context)
}

type Messenger interface {
	GetChatsHandler(c *gin.Context)
	SendMessageHandler(c *gin.Context)
	GetChatHandler(c *gin.Context)
}

type Services struct {
	Authentication Authentication
	Messenger      Messenger
}

func NewServices(repos *repository.Repositories, tokenManager auth.TokenManager) *Services {
	authenticationService := NewAuthenticationService(repos.Authentication, tokenManager)
	messengerService := NewMessengerService(repos.Messenger)

	return &Services{
		authenticationService,
		messengerService,
	}
}
