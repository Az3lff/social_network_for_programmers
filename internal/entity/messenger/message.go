package entity

type Message struct {
	Id string `db:"message_id" json"id"`
	Content string `db:"content" json:"content"`
	// Username string `json:"Username"`
}
