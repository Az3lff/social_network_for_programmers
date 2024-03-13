package Messanger

import (

	"github.com/gin-gonic/gin"
)


var users []User = []User{User{Id: 1, Name: "Kirill"}, User{Id: 2,Name: "Sergey"}}
// var messages []Message = []Message{Content: "хуй", chatid: 1}
var chats []Chat = []Chat{Chat{Id: 1, Profile: users[0], }}
// var Users = map[int]string{1 : "Kirill", 2: "Sergey"}
// var messages = map[int][]string{1: {"Hello peidor", "Hi"}, 2: {"yapidor", "tin ety"}}

func (h *handler) GetChatsHandler(c *gin.Context){
	// c.Writer.Write([]byte(fmt.Sprintf("%s", Users)))
}


func (h *handler) GetChatHandler(c *gin.Context){
	// chatid, _ := strconv.Atoi(c.Params.ByName("UserId")) 
	// c.Writer.Write([]byte(messages[chatid][len(messages[chatid])-1]))
}


func (h *handler) SendMessageHandler(c *gin.Context){

}