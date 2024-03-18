package entity

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	Chat    *Chat
	Message chan []byte
	Socket  *websocket.Conn
}

