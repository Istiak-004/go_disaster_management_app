package controllers

import (
	"net/http"
	"github.com/disaster_management_backend/internals/models"
	"github.com/disaster_management_backend/internals/database"
	"github.com/gin-gonic/gin"
)

func ReportCrisis(c *gin.Context) {
	var crisis models.Crisis
	if err := c.ShouldBindJSON(&crisis); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	crisis.Status = "pending"
	if err := database.DB.Create(&crisis).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not report crisis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Crisis reported successfully"})
}

func ApproveCrisis(c *gin.Context) {
	var crisis models.Crisis
	if err := c.ShouldBindJSON(&crisis); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := database.DB.Model(&crisis).Where("id = ?", crisis.ID).Update("status", "approved").Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not approve crisis"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "crisis approved!"})
}



type RespondToCrisisRequest struct {
    CrisisID uint   `json:"crisis_id" binding:"required"`
    Response string `json:"response" binding:"required"`
}

func RespondToCrisis(c *gin.Context) {
    var input RespondToCrisisRequest

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var crisis models.Crisis
    if err := database.DB.Where("id = ?", input.CrisisID).First(&crisis).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Crisis not found"})
        return
    }

    volunteerID := c.GetUint("userId")

    response := models.CrisisResponse{
        VolunteerID: volunteerID,
        CrisisID:    input.CrisisID,
        Response:    input.Response,
    }

    if err := database.DB.Create(&response).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save response"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Response recorded successfully", "response": response})
}