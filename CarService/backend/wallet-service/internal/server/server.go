package server

import (
	"log"
	"wallet-service/internal/database"

	"github.com/gin-gonic/gin"
)

func Start() {
	s := gin.Default()

	database.ConnectDB()

	err := s.Run(":8081")
	if err != nil {
		log.Fatal("Failed to create server")
	}
}
