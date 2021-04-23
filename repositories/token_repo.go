package repositories

import (
	"github.com/ducnt114/go-gin-demo/models"
	"github.com/jinzhu/gorm"
)

type TokenRepository interface {
	Save(m *models.Token) error
	Delete(m *models.Token) error
	FindByRefreshToken(refreshToken string) (*models.Token, error)
}

type tokenRepoImpl struct {
	orm *gorm.DB
}

func newTokenRepository(orm *gorm.DB) TokenRepository {
	return &tokenRepoImpl{
		orm: orm,
	}
}

func (r *tokenRepoImpl) Save(m *models.Token) error {
	return r.orm.Create(m).Error
}

func (r *tokenRepoImpl) Delete(m *models.Token) error {
	return r.orm.Delete(m).Error
}

func (r *tokenRepoImpl) FindByRefreshToken(refreshToken string) (*models.Token, error) {
	var res models.Token
	err := r.orm.Model(&models.User{}).Where("refresh_token = ?", refreshToken).First(&res).Error
	return &res, err
}
