package server

import (
	"CarService/internal/database"
	"CarService/internal/handlers"
	"log"

	"github.com/gin-gonic/gin"
)

func Start() {
	s := gin.Default()

	database.ConnectDB()

	s.Static("/frontend", "/Users/user/important/College/CarService/frontend")

	s.GET("/", handlers.LoadFrontend)

	err := s.Run(":8080")
	if err != nil {
		log.Fatal("Failed to create server")
	}
}
