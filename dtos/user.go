package dtos

type GetUserInfoResponse struct {
	Meta *Meta     `json:"meta"`
	Data *UserInfo `json:"data"`
}

type UserInfo struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"user_name"`
}
