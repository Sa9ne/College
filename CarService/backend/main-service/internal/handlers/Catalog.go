package handlers

import (
	"CarService/internal/database"
	"CarService/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Catalog(ctx *gin.Context) {
	var Cars []models.Car
	if err := database.DB.Find(&Cars).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Cars": Cars})
}
