package repositories

import (
	"github.com/ducnt114/go-gin-demo/models"
	"github.com/jinzhu/gorm"
)

type UserRepository interface {
	FindByID(userID uint) (*models.User, error)
	FindByName(name string) (*models.User, error)
}

type userRepoImpl struct {
	orm *gorm.DB
}

func newUserRepository(orm *gorm.DB) UserRepository {
	return &userRepoImpl{
		orm: orm,
	}
}

func (r *userRepoImpl) FindByID(userID uint) (*models.User, error) {
	var res models.User
	err := r.orm.Model(&models.User{}).Where("id = ?", userID).First(&res).Error
	return &res, err
}

func (r *userRepoImpl) FindByName(name string) (*models.User, error) {
	var res models.User
	err := r.orm.Model(&models.User{}).Where("name = ?", name).First(&res).Error
	return &res, err
}
