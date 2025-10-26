package db

import (
	"github.com/Kartik30R/Tiket.git/models"
	"gorm.io/gorm"
)

func DBMigrator(db *gorm.DB) error{
	return db.AutoMigrate(&models.Event{},&models.Ticket{}, &models.User{})
}