package v1

import "github.com/gin-gonic/gin"

func (h *Handler) initMessengerRoutes(api *gin.RouterGroup) {
	messenger := api.Group("/messenger")
	{
		// messenger.GET("/page", func(c *gin.Context) {
		//		// 	c.HTML(http.StatusOK, "test.html", gin.H{})
		//		// })

		messenger.POST("/", h.CreateChatHandler)
		messenger.GET("/chats/:UserId", h.GetAllChatsHandler)
		messenger.GET("/chat/:ChatId", h.GetChatHandler)
		messenger.GET("/ws/:chatId", h.GetConnChatHandler)
	}
}

func (h *Handler) CreateChatHandler(c *gin.Context) {

}

func (h *Handler) GetAllChatsHandler(c *gin.Context) {

}

func (h *Handler) GetChatHandler(c *gin.Context) {

}

func (h *Handler) GetConnChatHandler(c *gin.Context) {

}
