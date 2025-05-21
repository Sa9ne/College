package handlers

import (
	"net/http"
	"wallet-service/internal/database"
	"wallet-service/internal/models"

	"github.com/gin-gonic/gin"
)

func GetCar(ctx *gin.Context) {
	// Get user
	UserValue, exists := ctx.Get("user")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	// Bring it into the required format
	user, ok := UserValue.(models.Users)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User not found in context"})
		return
	}

	// Find cars
	var orders []models.Orders
	if err := database.DB.Where("user_id = ?", user.Id).Find(&orders).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Nothing found"})
		return
	}

	if len(orders) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No one car was bought"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"User have:": orders})
}
