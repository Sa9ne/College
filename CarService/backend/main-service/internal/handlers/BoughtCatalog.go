package handlers

import (
	"CarService/internal/database"
	"CarService/internal/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BoughtCatalog(ctx *gin.Context) {

	var BoughtCar []models.Orders
	if err := database.DB.Find(&BoughtCar).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Bought car was not found"})
		return
	}

	ctx.JSON(http.StatusOK, BoughtCar)
}
