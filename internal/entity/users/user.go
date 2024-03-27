package users

type UserSignUp struct {
	Login    string `json:"login"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserSignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRepo struct {
	UUID         string
	Login        string
	Email        string
	HashPassword string
}

type UserRestore struct {
	Email string `json:"email"`
}
