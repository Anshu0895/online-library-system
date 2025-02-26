package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"online-library-system/models"
	"testing"
)

// SetupTestDB sets up an in-memory test database
func SetupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}
	db.AutoMigrate(&models.Library{}, &models.User{}, &models.BookInventory{}, &models.RequestEvent{}, &models.IssueRegistry{})
	DB = db
	return db
}

// TestDatabaseConnection tests the database connection
func TestDatabaseConnection(t *testing.T) {
	db := SetupTestDB()
	if db == nil {
		t.Fatalf("Database connection is nil")
	}
}

// TestAutoMigrate tests the AutoMigrate function
func TestAutoMigrate(t *testing.T) {
	db := SetupTestDB()

	if !db.Migrator().HasTable(&models.Library{}) {
		t.Fatalf("Library table does not exist")
	}
	if !db.Migrator().HasTable(&models.User{}) {
		t.Fatalf("User table does not exist")
	}
	if !db.Migrator().HasTable(&models.BookInventory{}) {
		t.Fatalf("BookInventory table does not exist")
	}
	if !db.Migrator().HasTable(&models.RequestEvent{}) {
		t.Fatalf("RequestEvent table does not exist")
	}
	if !db.Migrator().HasTable(&models.IssueRegistry{}) {
		t.Fatalf("IssueRegistry table does not exist")
	}
}

// TestCRUDOperations tests basic CRUD operations
func TestCRUDOperations(t *testing.T) {
	db := SetupTestDB()

	// Create
	library := models.Library{ID: 1, Name: "Test Library"}
	result := db.Create(&library)
	if result.Error != nil {
		t.Fatalf("Failed to create library: %v", result.Error)
	}

	// Read
	var readLibrary models.Library
	result = db.First(&readLibrary, 1)
	if result.Error != nil {
		t.Fatalf("Failed to read library: %v", result.Error)
	}

	if readLibrary.Name != "Test Library" {
		t.Fatalf("Expected library name 'Test Library', got '%s'", readLibrary.Name)
	}

	// Update
	readLibrary.Name = "Updated Library"
	result = db.Save(&readLibrary)
	if result.Error != nil {
		t.Fatalf("Failed to update library: %v", result.Error)
	}

	var updatedLibrary models.Library
	db.First(&updatedLibrary, 1)
	if updatedLibrary.Name != "Updated Library" {
		t.Fatalf("Expected library name 'Updated Library', got '%s'", updatedLibrary.Name)
	}

	// Delete
	result = db.Delete(&models.Library{}, 1)
	if result.Error != nil {
		t.Fatalf("Failed to delete library: %v", result.Error)
	}

	var deletedLibrary models.Library
	result = db.First(&deletedLibrary, 1)
	if result.Error == nil {
		t.Fatalf("Expected no library, but found one")
	}
}
