package services

import (
	"net/http"
	"testing"

	mockRepo "github.com/ducnt114/go-gin-demo/mocks/repositories"
	mockUtils "github.com/ducnt114/go-gin-demo/mocks/utils"
	"github.com/ducnt114/go-gin-demo/models"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/mock"
)

func TestLogin(t *testing.T) {
	mockUserRepo := &mockRepo.MockUserRepo{}
	mockTokenRepo := &mockRepo.MockTokenRepo{}
	mockJWTHelper := &mockUtils.MockJWTHelper{}

	mockUserRepo.On("FindByName", "user-1").
		Return(nil, gorm.ErrRecordNotFound)
	mockUserRepo.On("FindByName", "user-2").
		Return(&models.User{
			Name:           "user-2",
			Salt:           "salt-2",
			HashedPassword: "hashed-pass-2",
		}, nil)
	mockUserRepo.On("FindByName", "user-3").
		Return(&models.User{
			Name:           "user-3",
			Salt:           "salt-3",
			HashedPassword: "b34593a8e1f8924bf57c72fb48c0f062c07938d574c364f7433dc89b7adb00a16536b1c069b1ecf690b390098baff2025779019785155e15d648fa4da8be1be3",
		}, nil)

	mockJWTHelper.On("GenerateToken", mock.Anything).Return("token-3", nil)

	mockTokenRepo.On("Save", mock.Anything).Return(nil)

	authService := newAuthService(mockUserRepo, mockTokenRepo, mockJWTHelper)

	// case not found
	res1 := authService.Login("user-1", "pass-1")
	if res1.Meta.Code != http.StatusNotFound {
		t.Fail()
	}

	// case hash password not match
	res2 := authService.Login("user-2", "pass-2")
	if res2.Meta.Code != http.StatusNotFound {
		t.Fail()
	}

	// case access token not match
	res3 := authService.Login("user-3", "pass-3")
	if res3.Data.AccessToken != "token-3" {
		t.Fail()
	}
}

func TestRefreshToken(t *testing.T) {
	mockUserRepo := &mockRepo.MockUserRepo{}
	mockTokenRepo := &mockRepo.MockTokenRepo{}
	mockJWTHelper := &mockUtils.MockJWTHelper{}

	mockUserRepo.On("FindByID", uint(2)).
		Return(nil, gorm.ErrRecordNotFound)
	mockUserRepo.On("FindByID", uint(3)).
		Return(&models.User{Name: "user-3"}, nil)

	mockTokenRepo.On("FindByRefreshToken", "old-token-1").Return(nil, gorm.ErrRecordNotFound)
	mockTokenRepo.On("FindByRefreshToken", "old-token-2").
		Return(&models.Token{UserID: uint(2), RefreshToken: "old-token-2"}, nil)
	mockTokenRepo.On("FindByRefreshToken", "old-token-3").
		Return(&models.Token{UserID: uint(3), RefreshToken: "old-token-3"}, nil)
	mockTokenRepo.On("Delete", mock.Anything).Return(nil)
	mockTokenRepo.On("Save", mock.Anything).Return(nil)

	mockJWTHelper.On("GenerateToken", mock.Anything).Return("token-3", nil)

	authService := newAuthService(mockUserRepo, mockTokenRepo, mockJWTHelper)

	// case refresh token not found
	res1 := authService.RefreshToken("old-token-1")
	if res1.Meta.Code != http.StatusNotFound {
		t.Fail()
	}

	// case user not found
	res2 := authService.RefreshToken("old-token-2")
	if res2.Meta.Code != http.StatusNotFound {
		t.Fail()
	}

	// case success
	res3 := authService.RefreshToken("old-token-3")
	if res3.Data.RefreshToken == "" {
		t.Fail()
	}
}
