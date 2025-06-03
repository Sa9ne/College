package handlers

import (
	"CarService/internal/database"
	"CarService/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Catalog(ctx *gin.Context) {
	var Cars []models.Car

	// Check bought status car and show result
	if err := database.DB.Where("bought_status = ?", false).Find(&Cars).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Not Found"})
		return
	}

	ctx.JSON(http.StatusOK, Cars)
}
