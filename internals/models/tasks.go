package models

import "gorm.io/gorm"

type Task struct {
    gorm.Model
    VolunteerID uint   `json:"volunteer_id"`
    Volunteer   User   `gorm:"foreignKey:VolunteerID"` 
    CrisisID    uint   `json:"crisis_id"`
    Crisis      Crisis `gorm:"foreignKey:CrisisID"` 
    Task        string `json:"task"`
    Status      string `json:"status"`
}
