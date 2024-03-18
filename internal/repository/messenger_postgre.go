package repository

import (
	"github.com/gin-gonic/gin"
	
)

type MessengerRepo struct {
}




func NewMessengerRepo() *MessengerRepo {
	return &MessengerRepo{}
}

func (m *MessengerRepo) GetChatsHandler(c *gin.Context) {
	
}

func (m *MessengerRepo) SendMessageHandler(c *gin.Context) {
	
}

func (m *MessengerRepo) GetChatHandler(c *gin.Context) {

}
