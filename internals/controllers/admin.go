package controllers

import (
	"net/http"

	"github.com/disaster_management_backend/internals/database"
	"github.com/disaster_management_backend/internals/models"
	"github.com/gin-gonic/gin"
)

type ApproveVolunteerRequest struct {
	VolunteerID uint `json:"volunteer_id" binding:"required"`
}

func ApproveVolunteer(c *gin.Context) {
	var input ApproveVolunteerRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var volunteer models.User
	if err := database.DB.Where("id = ? AND role = 'volunteer'", input.VolunteerID).First(&volunteer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "volunteer not found"})
		return
	}

	volunteer.IsVerified = true
	if err := database.DB.Save(&volunteer).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "could not approve volunteer!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Volunteer approved successfully!", "volunteer": volunteer})
}
