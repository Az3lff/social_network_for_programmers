package users

type RestoreAccessRequest struct {
	Email string `yaml:"email"`
	Code  string `yaml:"code"`
}
