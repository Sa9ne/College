package handlers

import (
	"net/http"
	"wallet-service/internal/database"
	"wallet-service/internal/models"

	"github.com/gin-gonic/gin"
)

func DeleteOrder(ctx *gin.Context) {
	ID := ctx.Param("id")

	var order models.Orders
	if err := database.DB.First(&order, ID).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	if err := database.DB.Delete(&order).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Order was not deleted"})
		return
	}

	if err := database.DB.Model(&models.Car{}).
		Where("vin = ?", order.VIN).
		Update("bought_status", false).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Bought status was not changed"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Order deleted successful"})
}
