package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email        string    `gorm:"uniqueIndex;not null"`
	Password     string    `gorm:"not null"`
	AccessToken  string    `gorm:"not null"`
	RefreshToken string    `gorm:"not null"`
	ExpiredAt    time.Time `gorm:"not null"`
}

func (User) TableName() string {
	return "users"
}
