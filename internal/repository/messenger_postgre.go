package repository

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type MessengerRepo struct {
	db *pgxpool.Conn
}

func NewMessengerRepo() *MessengerRepo {
	return &MessengerRepo{}
}

func (m *MessengerRepo) GetChatsHandler() {

}

func (m *MessengerRepo) SendMessageHandler() {

}

func (m *MessengerRepo) GetChatHandler() {

}
