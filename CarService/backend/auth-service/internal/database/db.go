package database

import (
	"auth-service/internal/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	// Read env file
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatalf("Config Database does't found")
	}

	// Database connection
	var errOpen error
	DB, errOpen = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errOpen != nil {
		log.Fatalf("Failed to connect DataBase: %v", errOpen)
	}

	DB.AutoMigrate(&models.Users{})
}
