package models

import (
	"gorm.io/gorm"
)

type Volunteer struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Tasks []Task `gorm:"foreignKey:VolunteerID"`
}
