package controllers

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/disaster_management_backend/internals/database"
	"github.com/disaster_management_backend/internals/models"
	"github.com/gin-gonic/gin"
)

// generates a CSV report for daily donations
func GenerateDailyDonationReport(c *gin.Context) {
	var donations []models.Donation
	today := time.Now().Format("2006-01-02")
	if err := database.DB.Where("DATE(created_at) = ?", today).Find(&donations).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch donation data"})
		return
	}

	file, err := os.Create("daily_donation_report.csv")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create CSV file"})
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Donor Name", "Amount", "Date"})

	for _, donation := range donations {
		writer.Write([]string{
			donation.DonorName,
			fmt.Sprintf("%.2f", donation.Amount),
			donation.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	c.File("daily_donation_report.csv")
}

// generates a CSV report for daily expenses
func GenerateDailyExpenseReport(c *gin.Context) {
	var expenses []models.Expense
	today := time.Now().Format("2006-01-02")
	if err := database.DB.Where("DATE(created_at) = ?", today).Find(&expenses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch expense data"})
		return
	}

	file, err := os.Create("daily_expense_report.csv")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create CSV file"})
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Item Name", "Amount", "Date"})

	for _, expense := range expenses {
		writer.Write([]string{
			expense.ItemName,
			fmt.Sprintf("%.2f", expense.Amount),
			expense.ExpenseBy,
			expense.CreatedAt.Format("2006-01-02 15:04:05"),
		})
	}

	c.File("daily_expense_report.csv")
}

// generates a CSV report for inventory items
func GenerateInventoryReportCSV(c *gin.Context) {
	var inventory []models.Inventory

	today := time.Now().Format("2006-01-02")
	if err := database.DB.Where("DATE(created_at) = ?", today).Find(&inventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch expense data"})
		return
	}

	file, err := os.Create("inventory_report.csv")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create CSV file"})
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"ID", "Item Name", "Quantity", "Price per Unit", "Purchased By"})

	for _, item := range inventory {
		writer.Write([]string{
			fmt.Sprintf("%d", item.ID),
			item.ItemName,
			fmt.Sprintf("%d", item.Quantity),
			fmt.Sprintf("%.2f", item.PricePerUnit),
			fmt.Sprintf("%d", item.PurchasedBy),
		})
	}

	c.File("inventory_report.csv")
}
