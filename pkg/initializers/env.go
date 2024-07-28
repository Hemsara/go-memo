package initializers

import (
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
)

func LoadENV() {
	rootDir := ".."
	envFilePath := filepath.Join(rootDir, ".env")
	if err := godotenv.Load(envFilePath); err != nil {
		log.Fatalf("Error loading %s file: %v", envFilePath, err)
	}

}
