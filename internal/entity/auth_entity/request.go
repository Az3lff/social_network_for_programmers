package auth_entity

type RestoreAccessRequest struct {
	Email string `yaml:"email"`
	Code  string `yaml:"code"`
}

type UpdatePasswordRequest struct {
	Email     string `yaml:"email"`
	Password1 string `yaml:"password1"`
	Password2 string `yaml:"password2"`
}
