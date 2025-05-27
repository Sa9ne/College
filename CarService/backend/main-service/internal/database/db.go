package database

import (
	"CarService/internal/models"
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

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("Database_url not found")
	}

	var errOpen error
	DB, errOpen = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errOpen != nil {
		log.Fatalf("Failed connect database: %v", errOpen)
	}

	DB.AutoMigrate(&models.Client{}, &models.Car{}, &models.Orders{})
}
