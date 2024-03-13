package Messanger

import (
	"time"

)


type User struct{
	Id int
	Name string
}

type Message struct {
	Content  string
	TimeSend time.Time 
	ChatId int
}

type Chat struct{
	Id int
	Profile User
	Messages []Message
}