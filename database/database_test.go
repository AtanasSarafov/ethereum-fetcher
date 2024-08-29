package database

import (
	"eth-fetcher/models"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// InitTestDB initializes an in-memory SQLite database for testing.
func InitTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the in-memory database!")
	}
	db.AutoMigrate(&models.Transaction{})
	return db
}

// TestInitDB tests the InitDB function to ensure it initializes the database correctly.
func TestInitDB(t *testing.T) {
	// Mock the gorm.Open function to return an in-memory SQLite database instead of a real PostgreSQL database
	db := InitTestDB()

	// Ensure the DB is not nil
	if db == nil {
		t.Fatal("Expected non-nil DB connection")
	}

	// Verify the schema migration (check if the 'transactions' table exists)
	if !db.Migrator().HasTable(&models.Transaction{}) {
		t.Error("Expected 'transactions' table to be created")
	}

	// You can add additional tests here, such as inserting a record and retrieving it to verify the database interaction
	var count int64
	db.Model(&models.Transaction{}).Count(&count)
	if count != 0 {
		t.Errorf("Expected 0 records in 'transactions' table, got %d", count)
	}
}
