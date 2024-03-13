package models

type User struct {
	ID       int    `json:"id"`
	NickName string `json:"nick_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
