package entities

type LoginErrorResponse struct {
	Error string `json:"error"`
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
