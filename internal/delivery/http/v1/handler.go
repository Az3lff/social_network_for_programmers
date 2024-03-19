package v1

import (
	"github.com/gin-gonic/gin"
	"social_network_for_programmers/internal/service"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	//Добавить главную страницу, на которой будет переход на регистрацию/авторизацию
	//router.GET("/soc_net_prog")

	messenger := router.Group("/messages")
	{
		messenger.GET("/", h.services.Messenger.GetChatsHandler)

		messages := messenger.Group("/:UserId")
		{
			messages.GET("/", h.services.Messenger.GetChatHandler)
			messages.POST("/:message", h.services.Messenger.SendMessageHandler)
		}
	}

	signUp := router.Group("/signUp")
	{
		signUp.GET("/", h.services.Authentication.SignUpPage)
		signUp.POST("/", h.services.Authentication.SignUp)
	}

	signIn := router.Group("/signIn")
	{
		signIn.GET("/", h.services.Authentication.SignInPage)
		signIn.POST("/", h.services.Authentication.SignIn)
	}

	return router
}
