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

	s.Static("/static", "/Users/user/important/College/CarService/frontend/build/web")

	s.GET("/", handlers.LoadFrontend)
	s.GET("/Catalog", handlers.Catalog)

	err := s.Run(":8080")
	if err != nil {
		log.Fatal("Failed to create server")
	}
}
