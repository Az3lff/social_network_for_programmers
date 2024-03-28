package responses

type ErrorResponse struct {
	Message string `json:"message"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}

type InfoResponse struct {
	Message string `yaml:"message"`
}
