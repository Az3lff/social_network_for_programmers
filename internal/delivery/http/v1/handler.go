package v1

import (
	// "net/http"
	"social_network_for_programmers/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()
	// router.LoadHTMLGlob("templates/*")

	//Добавить главную страницу, на которой будет переход на регистрацию/авторизацию
	//router.GET("/soc_net_prog")

	messenger := router.Group("/messages")
	{
		messenger.GET("/", h.services.Messenger.GetChatsHandler)
		// messenger.GET("/page", func(c *gin.Context) {
		// 	c.HTML(http.StatusOK, "test.html", gin.H{})
		// })

		messages := messenger.Group("/ws")
		{
			messages.GET("/chats/:id", h.services.Messenger.GetChatHandler)
			messages.GET("/:ChatId", h.services.Messenger.SendMessageHandler)
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
