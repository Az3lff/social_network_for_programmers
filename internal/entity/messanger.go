package entity

import (
	"time"
)

type ChatUser struct {
	Id   int
	Name string
}

type Message struct {
	Content  string
	TimeSend time.Time
}

type Chat struct {
	Id int
	//Profile  User
	Messages []Message
}
