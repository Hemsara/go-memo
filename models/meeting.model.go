package models

import (
	"time"

	"gorm.io/gorm"
)

type Meeting struct {
	gorm.Model
	MeetingID    string `gorm:"uniqueIndex;not null"`
	UserID       uint   `gorm:"not null"` // Foreign key linking to User
	User         User   `gorm:"foreignKey:UserID"`
	SpecialNotes string `gorm:"type:text"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (Meeting) TableName() string {
	return "meetings"
}
