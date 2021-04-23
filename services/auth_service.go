package services

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ducnt114/go-gin-demo/dtos"
	"github.com/ducnt114/go-gin-demo/repositories"
	"github.com/ducnt114/go-gin-demo/utils"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type AuthService interface {
	Login(username, rawPass string) *dtos.LoginResponse
}

type authServiceImpl struct {
	userRepo  repositories.UserRepository
	jwtHelper utils.JWTHelper
}

func newAuthService(userRepo repositories.UserRepository, jwtHelper utils.JWTHelper) AuthService {
	return &authServiceImpl{
		userRepo:  userRepo,
		jwtHelper: jwtHelper,
	}
}

func (s *authServiceImpl) Login(username, rawPass string) *dtos.LoginResponse {
	user, err := s.userRepo.FindByName(username)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.S().Warn("User not found")
			return &dtos.LoginResponse{
				Meta: &dtos.Meta{
					Code:    http.StatusNotFound,
					Message: "user/pass not correct",
				},
			}
		}
		zap.S().Error("error when find user by name, detail: ", err)
		return &dtos.LoginResponse{Meta: dtos.InternalServerErrorMeta}
	}
	// compare password
	if utils.HashPassword(rawPass, user.Salt) != user.HashedPassword {
		zap.S().Warn("Password not match")
		return &dtos.LoginResponse{
			Meta: &dtos.Meta{
				Code:    http.StatusNotFound,
				Message: "user/pass not correct",
			},
		}
	}

	// generate token
	currentTime := time.Now()
	claims := dtos.AuthClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: currentTime.Add(10 * time.Hour).Unix(),
			IssuedAt:  currentTime.Unix(),
		},
		UserID:   user.ID,
		UserName: user.Name,
	}

	accessToken, err := s.jwtHelper.GenerateToken(&claims)
	if err != nil {
		zap.S().Errorf("Generating token for user: %v, err: %v", user.Name, err)
		return &dtos.LoginResponse{Meta: dtos.InternalServerErrorMeta}
	}

	return &dtos.LoginResponse{
		Meta: dtos.SuccessMeta,
		Data: &dtos.LoginResponseData{
			AccessToken:  accessToken,
			RefreshToken: "",
		},
	}
}
