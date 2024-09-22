package models

import (
	"gorm.io/gorm"
)

type Crisis struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Location    string `json:"location"`
	Severity    int    `json:"severity"`
	Status      string `json:"status"` // Pending, Approved, Resolved
	ImageURL    string `json:"image_url"`
	HelpNeeded  string `json:"help_needed"`
}

type CrisisResponse struct {
    gorm.Model
    VolunteerID uint   `json:"volunteer_id"`
    Volunteer   User   `gorm:"foreignKey:VolunteerID"` 
    CrisisID    uint   `json:"crisis_id"`
    Crisis      Crisis `gorm:"foreignKey:CrisisID"`  
    Response    string `json:"response"`
}
