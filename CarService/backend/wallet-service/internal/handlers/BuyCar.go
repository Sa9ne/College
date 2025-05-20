package handlers

import (
	"net/http"
	"strconv"
	"wallet-service/internal/database"
	"wallet-service/internal/models"

	"github.com/gin-gonic/gin"
)

func BuyCar(ctx *gin.Context) {
	// Get user in context
	UserValue, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Bring it into the required format
	user, ok := UserValue.(models.Users)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user in context"})
		return
	}

	// Get car id
	CarID, err := strconv.Atoi(ctx.Param("CarID"))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Invalid car ID"})
		return
	}

	// Find car
	var car models.Car
	if err := database.DB.First(&car, CarID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Car not found"})
		return
	}

	// Check wallet
	if user.Wallet < car.Price {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Not enough money"})
		return
	}

	// Add car in user profile
	order := models.Orders{
		UserID: user.Id,
		CarID:  car.Id,
	}

	if err := database.DB.Create(&order).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save order"})
		return
	}

	// Write off money into user wallet
	user.Wallet -= car.Price

	if err := database.DB.Save(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update wallet"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successful order", "wallet": user.Wallet})
}
