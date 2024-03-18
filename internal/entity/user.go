package entity

type UsersSignUpInput struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UsersSignInInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
