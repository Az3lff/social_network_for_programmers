package repository

import (
	"context"
	"log"
	entity "social_network_for_programmers/internal/entity/messenger"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MessengerRepo struct {
	db *pgxpool.Pool
}

func NewMessengerRepo(database *pgxpool.Pool) *MessengerRepo {
	return &MessengerRepo{db: database}
}

func (m *MessengerRepo) GetAllChats(UserId string) ([]string, error) {
	row, err := m.db.Query(context.Background(), "select chat_id from chat_users where user_id=$1", UserId)
	if err != nil {
		return []string{}, err
	}
	chats, err := pgx.CollectRows(row, pgx.RowTo[string])
	if err != nil{
		return []string{}, err
	}
	return chats, nil
}

func (m *MessengerRepo) CreateChat(senderId string, recipientId string) (string, error) {
	check, err := m.db.Query(context.Background(), "SELECT chat_id FROM public.chat_users WHERE user_id in ($1, $2) GROUP BY chat_id HAVING COUNT(*)>1", senderId, recipientId)
	if err != nil {
		return "", err
	}
	if isCreated, err := pgx.CollectOneRow(check, pgx.RowTo[string]); isCreated != ""{
		if err != nil{
			return "", err
		}
		return isCreated, nil
	}
	row, err := m.db.Query(context.Background(), "insert into chats (chat_id) values (gen_random_uuid()) returning chat_id")
	if err != nil {
		return "", err
	}
	ChatId, err := pgx.CollectOneRow(row, pgx.RowTo[string])
	if err != nil {
		return "", err
	}
	for _, user := range []string{senderId, recipientId} {
		_, err := m.db.Query(context.Background(), "insert into chat_users (chat_id, user_id) values ($1, &2)", ChatId, user)
		if err != nil {
			return "", err
		}
	}
	return ChatId, nil
}

func (m *MessengerRepo) GetMessages(ChatId string, messages *[]entity.Message) error {
	rows, err := m.db.Query(context.Background(), "select message_id, content from messages where chat_id=$1", ChatId)
	if err != nil {
		return err
	}
	*messages, err = pgx.CollectRows(rows, pgx.RowToStructByName[entity.Message])
	if err != nil {
		return err
	}
	return nil
}

func (m *MessengerRepo) SaveMessage(ChatId string, mess *entity.Message) error {
	log.Println(mess.Content)
	_, err := m.db.Query(context.Background(), "INSERT INTO messages (content, chat_id) VALUES ($1, $2)", mess.Content, ChatId)
	if err != nil {
		return err
	}
	return nil
}
