package handlers

import (
	"net/http"
	"wallet-service/internal/database"
	"wallet-service/internal/models"

	"github.com/gin-gonic/gin"
)

func MakeOrder(ctx *gin.Context) {
	var req models.Orders

	// Parsing JSON
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var car models.Car

	// Checking win number car
	if err := database.DB.First(&car, "vin = ?", req.VIN).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Car was not found"})
		return
	}

	// Check if car was bought
	if car.BoughtStatus {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Car was bought"})
		return
	}

	var client models.Client

	// Check phone number
	if err := database.DB.First(&client, "phone = ?", req.Phone).Error; err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect phone request"})
		return
	}

	// Add order
	if err := database.DB.Create(&req).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	// Update car status
	if err := database.DB.Model(&car).Update("bought_status", true).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update car status"})
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Created order successful"})
}
