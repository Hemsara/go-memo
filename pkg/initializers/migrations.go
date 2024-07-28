package initializers

import (
	"calendar_automation/models"
	"calendar_automation/pkg/database"
	"fmt"
)

func MakeMigrations() {
	if database.DB == nil {
		panic("database connection is not initialized")
	}

	err := database.DB.AutoMigrate(
		&models.User{},
		&models.Meeting{},
	)
	if err != nil {
		panic(fmt.Sprintf("cannot auto-migrate database: %s", err))
	}
}
