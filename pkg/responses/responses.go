package responses

type ErrorBadRequest struct {
	Err string `json:"error"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
