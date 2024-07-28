package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DBUsername = "DB_USERNAME"
	DBPassword = "DB_PASSWORD"
	DBHost     = "DB_HOST"
	DBPort     = "DB_PORT"
	DBDatabase = "DB_DATABASE"
)

var DB *gorm.DB

func New() {
	if DB != nil {
		return
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv(DBHost),
		os.Getenv(DBUsername),
		os.Getenv(DBPassword),
		os.Getenv(DBDatabase),
		os.Getenv(DBPort),
	)

	var err error
	d, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	DB = d
}

func Close() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
