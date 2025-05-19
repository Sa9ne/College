package handlers

import (
	"net/http"
	"wallet-service/internal/database"
	"wallet-service/internal/models"

	"github.com/gin-gonic/gin"
)

type MoneyReq struct {
	Amount int `json:"amount"`
}

func MoneyAdd(ctx *gin.Context) {
	// Получаем пользователя
	userValue, exist := ctx.Get("user")
	if !exist {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not found in context"})
		return
	}

	user, ok := userValue.(models.Users)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user in context"})
		return
	}

	// Получаем информацию о деньгах
	var Money MoneyReq
	if err := ctx.ShouldBindJSON(&Money); err != nil || Money.Amount <= 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid amount"})
		return
	}

	// Добавляем в кошелек деньги
	user.Wallet += Money.Amount
	if err := database.DB.Save(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Money doesn't save"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Wallet": user.Wallet})
}
