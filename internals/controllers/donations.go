package controllers

import (
	"net/http"
	"github.com/disaster_management_backend/internals/models"
	"github.com/disaster_management_backend/internals/database"
	"github.com/gin-gonic/gin"
)

func AddDonation(c *gin.Context) {
	var donation models.Donation
	var transaction models.Transaction
	if err := c.ShouldBindJSON(&donation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Create(&donation).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not process donation"})
		return
	}

	if err:= database.DB.Create(&transaction).Where("transaction_type=?","donation").Error;err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create transaction record"})
        return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Donation successful"})
}

func GetTotalDonations(c *gin.Context) {
	var total float64
	database.DB.Model(&models.Donation{}).Select("SUM(amount)").Scan(&total)
	c.JSON(http.StatusOK, gin.H{"total_donations": total})
}
