package database

import (
	"eth-fetcher/config"
	"eth-fetcher/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.DatabaseDSN), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database!")
	}
	// Automatically migrate the schema
	db.AutoMigrate(&models.Transaction{})
	return db
}
