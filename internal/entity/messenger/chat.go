package entity

type Chat struct {
	Clients  map[*Client]bool
	Join     chan *Client
	Leave    chan *Client
	Messages chan []byte
}


