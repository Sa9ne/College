package database

import (
	"auth-service/internal/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	err := godotenv.Load("/Users/user/important/College/CarService/.env")
	if err != nil {
		log.Fatal("Failed to load .env file")
	}

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
