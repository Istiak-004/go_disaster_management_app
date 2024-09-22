package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string `json:"name"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Role       string `json:"role"` // Admin or Volunteer
	Password   string `json:"password"`
	IsVerified bool   `json:"is_verified"`
}
