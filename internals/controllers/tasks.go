package controllers

import (
    "github.com/disaster_management_backend/internals/database"
    "github.com/disaster_management_backend/internals/models"
    "github.com/gin-gonic/gin"
    "net/http"
)

type AssignTaskRequest struct {
    VolunteerID uint   `json:"volunteer_id" binding:"required"`
    CrisisID    uint   `json:"crisis_id" binding:"required"`
    Task        string `json:"task" binding:"required"`
}

func AssignTask(c *gin.Context) {
    var input AssignTaskRequest

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var volunteer models.User
    if err := database.DB.Where("id = ? AND role = ? AND is_verified=?", input.VolunteerID, "volunteer", true).First(&volunteer).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Volunteer not found"})
        return
    }

    var crisis models.Crisis
    if err := database.DB.Where("id = ?", input.CrisisID).First(&crisis).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Crisis not found"})
        return
    }

    taskAssignment := models.Task{
        VolunteerID: input.VolunteerID,
        CrisisID:    input.CrisisID,
        Task:        input.Task,
        Status:      "assigned", 
    }

    if err := database.DB.Create(&taskAssignment).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not assign task"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Task assigned successfully", "task": taskAssignment})
}
