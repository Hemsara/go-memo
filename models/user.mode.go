package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email          string    `gorm:"uniqueIndex;not null"`
	Password       string    `gorm:"not null"`
	FullName       string    `gorm:"not null"`
	AccessToken    string    `gorm:""`
	RefreshToken   string    `gorm:""`
	TokenExpiredAt time.Time `gorm:""`
}

func (User) TableName() string {
	return "users"
}
