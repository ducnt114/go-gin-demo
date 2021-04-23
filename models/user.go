package models

import "time"

type User struct {
	ID             int64     `gorm:"column:id;primary_key"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
	Name           string    `gorm:"column:name"`
	HashedPassword string    `gorm:"column:hashed_password"`
	Salt           string    `gorm:"column:salt"`
}

func (User) TableName() string {
	return "user"
}
