package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name           string `gorm:"column:name"`
	HashedPassword string `gorm:"column:hashed_password"`
	Salt           string `gorm:"column:salt"`
}

func (User) TableName() string {
	return "user"
}
