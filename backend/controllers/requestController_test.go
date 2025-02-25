package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"online-library-system/models"
	"testing"
)

// Setup test database

// Test Create Request
func TestCreateIssueRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/requests", RaiseIssueRequest)

	db := SetupTestDB()

	// Create a test user (reader)
	reader := models.User{
		Name:          "John Reader",
		Email:         "reader@example.com",
		Password:      "password123",
		ContactNumber: "9876543210",
		Role:          "Reader",
	}
	db.Create(&reader)

	// Create a test book
	book := models.BookInventory{
		ISBN:            "1234567890",
		Title:           "Test Book",
		Authors:         "Test Author",
		AvailableCopies: 2,
	}
	db.Create(&book)

	// Ensure book exists
	var fetchedBook models.BookInventory
	db.First(&fetchedBook, "isbn = ?", book.ISBN)
	if fetchedBook.ISBN == "" {
		t.Fatalf("Failed to fetch the book from database")
	}

	// Create request payload
	requestData := map[string]interface{}{
		"book_id":   fetchedBook.ISBN, // Use fetched book ID
		"reader_id": reader.ID,
	}

	body, _ := json.Marshal(requestData)
	req, _ := http.NewRequest("POST", "/requests", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	// Fetch the request from database
	var createdRequest models.RequestEvent
	db.First(&createdRequest)

	if createdRequest.BookID != fetchedBook.ISBN {
		t.Errorf("handler returned unexpected book_id: got %v want %v", createdRequest.BookID, fetchedBook.ISBN)
	}
}
