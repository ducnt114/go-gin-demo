package services

import (
	"net/http"

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
	userRepo repositories.UserRepository
}

func newAuthService(userRepo repositories.UserRepository) AuthService {
	return &authServiceImpl{
		userRepo: userRepo,
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

	return &dtos.LoginResponse{
		Meta: dtos.SuccessMeta,
		Data: &dtos.LoginResponseData{},
	}
}
