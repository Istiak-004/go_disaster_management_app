// controllers/chart.go
package controllers

import (
	"net/http"

	"github.com/disaster_management_backend/internals/database"
	"github.com/disaster_management_backend/internals/models"
	"github.com/gin-gonic/gin"
)

type CartsReportResponse struct {
	Donations []models.Donation `json:"donations"`
	Expenses  []models.Expense  `json:"expenses"`
}

func GetFundsAndExpenses(c *gin.Context) {
	var expenses []models.Expense
	var donations []models.Donation

	if err := database.DB.Find(&expenses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch expenses"})
		return
	}

	if err := database.DB.Find(&donations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch doantaions"})
		return
	}

	results := []CartsReportResponse{
		{
			Donations: donations,
			Expenses:  expenses,
		},
	}

	c.JSON(http.StatusOK, gin.H{"data": results})
}
