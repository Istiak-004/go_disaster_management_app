// database/db.go
package database

import (
	"log"

	"github.com/disaster_management_backend/internals/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(&models.User{}, &models.Donation{}, &models.Crisis{}, &models.Inventory{}, &models.Task{}, &models.Transaction{}, &models.Inventory{}, &models.Expense{})
	DB = db
	return db
}
