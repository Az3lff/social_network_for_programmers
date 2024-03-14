package entity

type User struct {
	ID       int    `json:"-"`
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
