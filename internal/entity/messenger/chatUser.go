package entity

type ChatUser struct {
	SenderId    string `json:"sender_id"`
	RecipientId string `json:"recipient_id"`
}
