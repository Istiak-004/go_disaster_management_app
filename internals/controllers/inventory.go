// controllers/inventory.go
package controllers

import (
	"net/http"

	"github.com/disaster_management_backend/internals/database"
	"github.com/disaster_management_backend/internals/models"
	"github.com/gin-gonic/gin"
)

func AddInventory(c *gin.Context) {
	var inventory models.Inventory
	if err := c.ShouldBindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&inventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to add inventory item!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inventory item added successfully!"})
}

func UpdateInventory(c *gin.Context) {
	var inventory models.Inventory
	id := c.Param("id")

	if err := database.DB.Where("id = ?", id).First(&inventory).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory item not found"})
		return
	}

	if err := c.ShouldBindJSON(&inventory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Save(&inventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update inventory item"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inventory item updated successfully!"})
}

func DeleteInventory(c *gin.Context) {
	id := c.Param("id")
	var inventory models.Inventory
	if err := database.DB.Where("id = ?", id).First(&inventory).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Inventory item not found"})
		return
	}

	if err := database.DB.Delete(&inventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete inventory item!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inventory item deleted successfully"})
}

func ListInventory(c *gin.Context) {
	var inventory []models.Inventory
	if err := database.DB.Find(&inventory).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve inventory items!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"inventory": inventory})
}
