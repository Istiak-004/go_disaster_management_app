package controllers

import (
	"net/http"
	"strconv"

	"github.com/disaster_management_backend/internals/database"
	"github.com/disaster_management_backend/internals/models"

	"github.com/gin-gonic/gin"
)

func GetVolunteerByID(c *gin.Context) {
	idParam := c.Param("id")
	volunteerID, err := strconv.ParseUint(idParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid volunteer ID"})
		return
	}

	var volunteer models.User
	if err := database.DB.Where("id = ? AND role = ?", volunteerID, "volunteer").First(&volunteer).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "volunteer not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"volunteer": volunteer})
}

type Volunteer struct {
	Name       string        `json:"name"`
	Email      string        `json:"email"`
	Phone      string        `json:"phone"`
	Tasks      []models.Task `json:"tasks"`
	IsVerified bool          `json:"is_verified"`
}

func GetAllVolunteer(c *gin.Context) {
	var volunteers []models.User
	var response []Volunteer
	if err := database.DB.Where("role = ?", "volunteer").Find(&volunteers).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "volunteer not found"})
		return
	}

	for _, volunteer := range volunteers {
		var tasks []models.Task
		_ = database.DB.Where("volunteer_id = ?", volunteer.ID).Find(&tasks)

		response = append(response, Volunteer{
			Name:       volunteer.Name,
			Email:      volunteer.Email,
			Phone:      volunteer.Phone,
			IsVerified: volunteer.IsVerified,
			Tasks:      tasks,
		})
	}

	c.JSON(http.StatusOK, gin.H{"volunteers": response})
}
