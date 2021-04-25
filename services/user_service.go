package services

import (
	"net/http"

	"github.com/ducnt114/go-gin-demo/dtos"
	"github.com/ducnt114/go-gin-demo/repositories"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

type UserService interface {
	GetUserInfo(userID uint) *dtos.GetUserInfoResponse
}

type userServiceImpl struct {
	userRepo repositories.UserRepository
}

func newUserService(userRepo repositories.UserRepository) UserService {
	return &userServiceImpl{
		userRepo: userRepo,
	}
}

func (r *userServiceImpl) GetUserInfo(userID uint) *dtos.GetUserInfoResponse {
	user, err := r.userRepo.FindByID(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			zap.S().Warn("User not found")
			return &dtos.GetUserInfoResponse{
				Meta: &dtos.Meta{
					Code:    http.StatusNotFound,
					Message: "user not found",
				},
			}
		}
		zap.S().Error("error when find user by id, detail: ", err)
		return &dtos.GetUserInfoResponse{Meta: dtos.InternalServerErrorMeta}
	}
	return &dtos.GetUserInfoResponse{
		Meta: &dtos.Meta{
			Code:    http.StatusOK,
			Message: "success",
		},
		Data: &dtos.UserInfo{
			UserID:   user.ID,
			Username: user.Name,
		},
	}
}
