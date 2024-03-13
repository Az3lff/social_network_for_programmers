package Messanger

import "time"

type Message struct {
	Content  string
	TimeSend time.Time
	//UserId int
}
