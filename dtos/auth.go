package dtos

import "github.com/dgrijalva/jwt-go"

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Meta *Meta              `json:"meta"`
	Data *LoginResponseData `json:"data"`
}

type LoginResponseData struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// AuthClaims represents jwt claims information.
type AuthClaims struct {
	jwt.StandardClaims
	UserID   uint   `json:"user_id,omitempty"`
	UserName string `json:"user_name,omitempty"`
}

type RefreshTokenReq struct {
	RefreshToken string `json:"refresh_token"`
}
