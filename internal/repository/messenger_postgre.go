package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MessengerRepo struct {
	db *pgxpool.Conn
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
