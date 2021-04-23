package dtos

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Meta *Meta              `json:"meta"`
	Data *LoginResponseData `json:"data"`
}

type LoginResponseData struct {
}
