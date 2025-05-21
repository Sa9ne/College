package server

import (
	"log"
	"wallet-service/internal/database"
	"wallet-service/internal/handlers"
	"wallet-service/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Start() {
	s := gin.Default()

	database.ConnectDB()

	s.Use(middleware.RequireAuth)

	s.POST("/MoneyAdd", handlers.MoneyAdd)

	Car := s.Group("/Car")
	Car.Use(middleware.RequireAuth)

	Car.GET("/Get", handlers.GetCar)
	Car.POST("/Buy/:CarID", handlers.BuyCar)

	err := s.Run(":8082")
	if err != nil {
		log.Fatal("Failed to create server")
	}
}
