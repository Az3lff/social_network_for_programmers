package service

import (
	"github.com/gin-gonic/gin"
	"social_network_for_programmers/internal/repository"
)

type MessengerService struct {
	repo repository.Messenger
}

func NewMessengerService(repo repository.Messenger) *MessengerService {
	return &MessengerService{repo}
}

func (m *MessengerService) GetChatsHandler(c *gin.Context) {

}

func (m *MessengerService) SendMessageHandler(c *gin.Context) {

}

func (m *MessengerService) GetChatHandler(c *gin.Context) {

}

//var users []entity.ChatUser = []entity.ChatUser{entity.ChatUser{Id: 1, Name: "Kirill"}, entity.ChatUser{Id: 2, Name: "Sergey"}}
//
//// var messages []Message = []Message{Content: "хуй", chatid: 1}
////var chats []entity.ChatUser = []entity.ChatUser{entity.ChatUser{Id: 1, Profile: users[0]}}
//
//// var Users = map[int]string{1 : "Kirill", 2: "Sergey"}
//// var messages = map[int][]string{1: {"Hello peidor", "Hi"}, 2: {"yapidor", "tin ety"}}
//
////func (h *delivery) GetChatsHandler(c *gin.Context) {
////	// c.Writer.Write([]byte(fmt.Sprintf("%s", Users)))
////}
////
////func (h v1.Handler) GetChatHandler(c *gin.Context) {
////	// chatid, _ := strconv.Atoi(c.Params.ByName("UserId"))
////	// c.Writer.Write([]byte(messages[chatid][len(messages[chatid])-1]))
////}
////
////func (h *handler) SendMessageHandler(c *gin.Context) {
////
////}