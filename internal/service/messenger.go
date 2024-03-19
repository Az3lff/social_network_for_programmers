package service

import (
	"log"
	"net/http"
	"strconv"
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

func (m *MessengerService) GetChatsHandler(c *gin.Context) {
	c.Writer.Write([]byte(""))
}

var upgrader = &websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024, CheckOrigin: func(r *http.Request) bool { return true }}

var rooms = make(map[int]map[*websocket.Conn]bool)

func (m *MessengerService) SendMessageHandler(c *gin.Context) {
	socket, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer socket.Close()
	chatId, err := strconv.Atoi(c.Params.ByName("ChatId")) 
	if err != nil{
		log.Println(err)
	}
	_, ok := rooms[chatId]

	if ok == false{
		rooms[chatId] = make(map[*websocket.Conn]bool)
	}
	rooms[chatId][socket] = true
	for {
		messageType, p, err := socket.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(p))
		for client := range rooms[chatId]{
			if err := client.WriteMessage(messageType, p); err != nil {
				log.Println(err)
				return
			}
		}
		
	}

}


var chats = map[int][]map[string]string{1: /* []map[string]string */{{"Kirill" : "message"}, {"Arseniy": "message"}}}
/* 
	{
		1: [
			{Name: message},

		]
	}
*/


func (m *MessengerService) GetChatHandler(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil{
		log.Println(err)
		return
	}
	c.JSON(http.StatusOK, chats[id])

}
