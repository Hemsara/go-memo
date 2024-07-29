package models

import "gorm.io/gorm"

type UserSession struct {
	gorm.Model

	ID uint `gorm:"primaryKey"` // Primary key field

	Country string `json:"country"`

	RegionName string `json:"regionName"`
	City       string `json:"city"`
	Zip        string `json:"zip"`

	Timezone string `json:"timezone"`
	ISP      string `json:"isp"`
	Query    string `json:"query"`
	UserID   uint   `gorm:"not null"`
	User     User   `gorm:"foreignKey:UserID"`
}

func (UserSession) TableName() string {
	return "sessions"
}
