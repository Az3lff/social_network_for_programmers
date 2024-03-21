package service

import (
	"log"
	"net/http"
	entity "social_network_for_programmers/internal/entity/messenger"
	"social_network_for_programmers/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type MessengerService struct {
	repo repository.Messenger
}

func NewMessengerService(repo repository.Messenger) *MessengerService {
	return &MessengerService{repo}
}

func (m *MessengerService) CreateChatHandler(c *gin.Context) {
	chatUsers := &entity.ChatUser{}
	err := c.BindJSON(&chatUsers)
	if err != nil{
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	chat_id, err := m.repo.CreateChat(chatUsers.SenderId, chatUsers.RecipientId)
	if err != nil{
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{"chat_id":chat_id})
}

var upgrader = &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024, CheckOrigin: func(r *http.Request) bool { return true }}

var rooms = make(map[string]map[*websocket.Conn]bool)

func (m *MessengerService) GetConnChatHandler(c *gin.Context) {
	socket, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	defer socket.Close()
	chatId := c.Params.ByName("ChatId")
	if _, ok := rooms[chatId];!ok {
		rooms[chatId] = make(map[*websocket.Conn]bool)
	}
	rooms[chatId][socket] = true
	for {
		messageType, p, err := socket.ReadMessage()
		if err != nil {
			log.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{})
			return
		}
		for client := range rooms[chatId] {
			if err := client.WriteMessage(messageType, p); err != nil {
				log.Println(err)
				c.JSON(http.StatusBadRequest, gin.H{})
				return
			}
			err = m.repo.SaveMessage(chatId, &entity.Message{Content: string(p)/* , Username: "" */})
			if err != nil{
				log.Println(err)
				c.JSON(http.StatusBadRequest, gin.H{})
				return
			}
		}
	}
}

func (m *MessengerService) GetChatHandler(c *gin.Context) {
	chatId := c.Params.ByName("ChatId")
	var messages []entity.Message
	err := m.repo.GetMessages(chatId, &messages)
	if err != nil{
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{chatId:messages})
}

func (m *MessengerService) GetAllChatsHandler(c *gin.Context){
	user_id := c.Params.ByName("UserId")
	chats, err := m.repo.GetAllChats(user_id)
	if err != nil{
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}
	c.JSON(http.StatusOK, gin.H{"chats_id":chats})
}