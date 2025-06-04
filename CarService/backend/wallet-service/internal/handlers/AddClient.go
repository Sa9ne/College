package handlers

import (
	"net/http"
	"regexp"
	"time"
	"wallet-service/internal/database"
	"wallet-service/internal/models"

	"github.com/gin-gonic/gin"
)

func isValidPhone(phone string) bool {
	re := regexp.MustCompile(`^(?:\+7|8)\d{10}$`)
	return re.MatchString(phone)
}

func AddClient(ctx *gin.Context) {
	var req models.Client

	// Parsing JSON
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check phone number
	if !isValidPhone(req.Phone) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Wrong phone number"})
		return
	}

	// Check birthday
	if req.Birthdate.Year() < 1900 || req.Birthdate.After(time.Now()) {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid birthdate"})
		return
	}

	// Create Client
	if err := database.DB.Create(&req).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "User was not created"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": req.ID})
}
