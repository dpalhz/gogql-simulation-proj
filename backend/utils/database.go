package utils

import (
	"log"
	"notes/backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitDB initializes the database connection and performs auto-migration.
func InitDB() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=gogql port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	autoMigrate := true

	if autoMigrate {
		if err := db.AutoMigrate(&models.User{}, &models.Note{}); err != nil {
			return nil, err
		}
		log.Println("Database migrated successfully!")
	}

	return db, nil
}
