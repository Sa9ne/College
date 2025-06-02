package server

import (
	"log"
	"wallet-service/internal/database"
	"wallet-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func Start() {
	s := gin.Default()

	database.ConnectDB()

	s.POST("/MakeOrder", handlers.MakeOrder)
	s.POST("/AddClient", handlers.AddClient)

	err := s.Run(":8081")
	if err != nil {
		log.Fatal("Failed to create server")
	}
}
