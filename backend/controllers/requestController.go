package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"online-library-system/database"
	"online-library-system/models"
	"time"
)

func RaiseIssueRequest(c *gin.Context) {
	var request models.RequestEvent
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Debugging: Print received reader ID
	fmt.Println("Received reader_id:", request.ReaderID)

	// Check if book exists
	var book models.BookInventory
	if err := database.DB.First(&book, "isbn = ?", request.BookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Check if book is available
	if book.AvailableCopies <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not available"})
		return
	}

	// Store request with correct reader ID
	request.RequestDate = time.Now()
	request.RequestType = "pending"
	if err := database.DB.Create(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Issue request raised", "request": request})
}

func GetRequestEvents(c *gin.Context) {
	var requestEvents []models.RequestEvent
	database.DB.Find(&requestEvents)
	c.JSON(http.StatusOK, requestEvents)
}

func GetRequestEventsByID(c *gin.Context) {
	reqID := c.Param("id")
	var request models.RequestEvent
	if err := database.DB.First(&request, "req_id = ?", reqID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Request not found"})
		return
	}
	c.JSON(http.StatusOK, request)
}

// Get Pending Requests
func GetPendingRequests(c *gin.Context) {
	var requests []models.RequestEvent

	if err := database.DB.Where("request_type = ?", "pending").Find(&requests).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"requests": requests})
}

func ApproveIssueRequest(c *gin.Context) {
	reqID := c.Param("id")
	var request models.RequestEvent

	// Get Request Details
	if err := database.DB.First(&request, "req_id = ?", reqID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Request not found"})
		return
	}

	// Check Book Availability
	var book models.BookInventory
	if err := database.DB.First(&book, "isbn = ?", request.BookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	if book.AvailableCopies <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not available"})
		return
	}

	// Approve Request - Expect JSON Body
	var reqBody struct {
		ApproverID uint `json:"approver_id"`
	}
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	log.Printf("Received ApproverId: %d", reqBody.ApproverID)

	request.ApprovalDate = time.Now()
	request.ApproverID = reqBody.ApproverID
	request.RequestType = "approved"

	if err := database.DB.Save(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Update Book Availability and Create Issue Registry
	book.AvailableCopies--
	if err := database.DB.Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	issue := models.IssueRegistry{
		ISBN:               request.BookID,
		ReaderID:           request.ReaderID,
		IssueApproverID:    request.ApproverID,
		IssueStatus:        "Issued",
		IssueDate:          time.Now(),
		ExpectedReturnDate: time.Now().AddDate(0, 0, 14), // 2 weeks from now
	}
	if err := database.DB.Create(&issue).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Issue request approved", "issue": issue})
}

func RejectIssueRequest(c *gin.Context) {
	reqID := c.Param("id")
	var request models.RequestEvent

	// Get Request Details
	if err := database.DB.First(&request, "req_id = ?", reqID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Request not found"})
		return
	}

	// Reject Request
	request.ApprovalDate = time.Now()
	request.ApproverID = 3

	request.RequestType = "rejected"
	if err := database.DB.Save(&request).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Issue request rejected"})
}
