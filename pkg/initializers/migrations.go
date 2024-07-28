package initializers

import (
	"calendar_automation/models"
	"calendar_automation/pkg/database"
	"fmt"
)

func MakeMigrations() {
	dbService := database.New()

	err := dbService.GetDB().AutoMigrate(

		&models.User{},
	)
	if err != nil {
		panic(fmt.Sprintf("cannot auto-migrate database: %s", err))
	}
}
