package controllers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"online-library-system/models"
	"testing"
	"time"
)

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

// TestGetPendingRequests tests the GetPendingRequests function
func TestGetPendingRequests(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/requests/pending", GetPendingRequests)

	db := SetupTestDB()

	// Create a test request event with "pending" type
	pendingRequest := models.RequestEvent{
		BookID:      "1234567890",
		ReaderID:    1,
		RequestDate: time.Now(),
		RequestType: "pending",
	}
	db.Create(&pendingRequest)

	// Create a test request event with "approved" type
	approvedRequest := models.RequestEvent{
		BookID:      "1234567890",
		ReaderID:    1,
		RequestDate: time.Now(),
		RequestType: "approved",
	}
	db.Create(&approvedRequest)

	req, _ := http.NewRequest("GET", "/requests/pending", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response struct {
		Requests []models.RequestEvent `json:"requests"`
	}
	json.Unmarshal(rr.Body.Bytes(), &response)

	if len(response.Requests) != 1 {
		t.Errorf("handler returned unexpected number of requests: got %v want %v", len(response.Requests), 1)
	}

	if response.Requests[0].RequestType != "pending" {
		t.Errorf("handler returned unexpected request type: got %v want %v", response.Requests[0].RequestType, "pending")
	}
}

// TestApproveIssueRequest tests the ApproveIssueRequest function
func TestApproveIssueRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.PUT("/requests/:id/approve", ApproveIssueRequest)

	db := SetupTestDB()

	// Create a test user (admin)
	admin := models.User{
		Name:          "Admin User",
		Email:         "admin@example.com",
		Password:      "password123",
		ContactNumber: "9876543210",
		Role:          "Admin",
	}
	db.Create(&admin)

	// Create a test book
	book := models.BookInventory{
		ISBN:            "1234567890",
		Title:           "Test Book",
		Authors:         "Test Author",
		AvailableCopies: 2,
	}
	db.Create(&book)

	// Create a test request event
	request := models.RequestEvent{
		ReqID:       1,
		BookID:      book.ISBN,
		ReaderID:    1,
		RequestDate: time.Now(),
		RequestType: "pending",
	}
	db.Create(&request)

	// Create approval payload
	approvalData := map[string]interface{}{
		"approver_id": admin.ID,
	}

	body, _ := json.Marshal(approvalData)
	req, _ := http.NewRequest("PUT", "/requests/1/approve", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Fetch the updated request from database
	var updatedRequest models.RequestEvent
	db.First(&updatedRequest, "req_id = ?", request.ReqID)

	if updatedRequest.RequestType != "approved" {
		t.Errorf("handler returned unexpected request type: got %v want %v", updatedRequest.RequestType, "approved")
	}

	if updatedRequest.ApproverID != admin.ID {
		t.Errorf("handler returned unexpected approver ID: got %v want %v", updatedRequest.ApproverID, admin.ID)
	}

	// Fetch the updated book from database
	var updatedBook models.BookInventory
	db.First(&updatedBook, "isbn = ?", book.ISBN)

	if updatedBook.AvailableCopies != book.AvailableCopies-1 {
		t.Errorf("handler returned unexpected available copies: got %v want %v", updatedBook.AvailableCopies, book.AvailableCopies-1)
	}

	// Fetch the issue record from database
	var issue models.IssueRegistry
	db.First(&issue, "isbn = ? AND reader_id = ?", book.ISBN, request.ReaderID)

	if issue.IssueStatus != "Issued" {
		t.Errorf("handler returned unexpected issue status: got %v want %v", issue.IssueStatus, "Issued")
	}

	if issue.IssueApproverID != admin.ID {
		t.Errorf("handler returned unexpected issue approver ID: got %v want %v", issue.IssueApproverID, admin.ID)
	}
}

// TestRejectIssueRequest tests the RejectIssueRequest function
func TestRejectIssueRequest(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.PUT("/requests/:id/reject", RejectIssueRequest)

	db := SetupTestDB()

	// Create a test user (admin)
	admin := models.User{
		Name:          "Admin User",
		Email:         "admin@example.com",
		Password:      "password123",
		ContactNumber: "9876543210",
		Role:          "Admin",
	}
	db.Create(&admin)

	// Create a test book
	book := models.BookInventory{
		ISBN:            "1234567890",
		Title:           "Test Book",
		Authors:         "Test Author",
		AvailableCopies: 2,
	}
	db.Create(&book)

	// Create a test request event
	request := models.RequestEvent{
		ReqID:       1,
		BookID:      book.ISBN,
		ReaderID:    1,
		RequestDate: time.Now(),
		RequestType: "pending",
	}
	db.Create(&request)

	// Create reject request payload
	rejectData := map[string]interface{}{
		"approver_id": admin.ID,
	}

	body, _ := json.Marshal(rejectData)
	req, _ := http.NewRequest("PUT", "/requests/1/reject", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Fetch the updated request from database
	var updatedRequest models.RequestEvent
	db.First(&updatedRequest, "req_id = ?", request.ReqID)

	if updatedRequest.RequestType != "rejected" {
		t.Errorf("handler returned unexpected request type: got %v want %v", updatedRequest.RequestType, "rejected")
	}

	if updatedRequest.ApproverID != 3 {
		t.Errorf("handler returned unexpected approver ID: got %v want %v", updatedRequest.ApproverID, 3)
	}
}
