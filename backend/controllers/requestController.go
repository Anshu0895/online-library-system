package controllers

import (
	"fmt"
	"log"
	"net/http"
	"online-library-system/database"
	"online-library-system/models"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Raise an issue request
// @Description Raise an issue request for a book
// @Tags requests
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param request body models.RequestEvent true "Request Data"
// @Success 201 {object} object "{"message": "Issue request raised", "request": models.RequestEvent}"
// @Failure 400 {object} object "{"error": "error message"}"
// @Failure 404 {object} object "{"error": "Book not found"}"
// @Failure 500 {object} object "{"error": "error message"}"
// @Router /requests/issue [post]

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

// @Summary Get all request events
// @Description Retrieve all request events
// @Tags requests
// @Accept json
// @Security ApiKeyAuth

// @Produce json
// @Success 200 {array} models.RequestEvent
// @Failure 500 {object} object "{"error": "error message"}"
// @Router /requests [get]
func GetRequestEvents(c *gin.Context) {
	var requestEvents []models.RequestEvent
	database.DB.Find(&requestEvents)
	c.JSON(http.StatusOK, requestEvents)
}

// @Summary Get request event by ID
// @Description Retrieve a request event by its ID
// @Tags requests
// @Accept json
// @Produce json
// @Security ApiKeyAuth

// @Param id path string true "Request ID"
// @Success 200 {object} models.RequestEvent
// @Failure 404 {object} object "{"error": "Request not found"}"
// @Failure 500 {object} object "{"error": "error message"}"
// @Router /requests/{id} [get]
func GetRequestEventsByID(c *gin.Context) {
	reqID := c.Param("id")
	var request models.RequestEvent
	if err := database.DB.First(&request, "req_id = ?", reqID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Request not found"})
		return
	}
	c.JSON(http.StatusOK, request)
}

// @Summary Get pending requests
// @Description Retrieve all pending requests
// @Tags requests
// @Accept json
// @Produce json
// @Security ApiKeyAuth

// @Success 200 {array} models.RequestEvent
// @Failure 500 {object} object "{"error": "error message"}"
// @Router /requests/pending [get]
func GetPendingRequests(c *gin.Context) {
	var requests []models.RequestEvent

	if err := database.DB.Where("request_type = ?", "pending").Find(&requests).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"requests": requests})
}

// @Summary Approve an issue request
// @Description Approve an issue request by ID
// @Tags requests
// @Accept json
// @Security ApiKeyAuth

// @Produce json
// @Param id path string true "Request ID"
// @Param approver body models.RequestEvent true "Approver Data"
// @Success 200 {object} object "{"message": "Issue request approved", "issue": models.IssueRegistry}"
// @Failure 400 {object} object "{"error": "Invalid request body"}"
// @Failure 404 {object} object "{"error": "Request not found"}"
// @Failure 500 {object} object "{"error": "error message"}"
// @Router /requests/{id}/approve [post]
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

// @Summary Reject an issue request
// @Description Reject an issue request by ID
// @Tags requests
// @Accept json
// @Produce json
// @Security ApiKeyAuth

// @Param id path string true "Request ID"
// @Success 200 {object} object "{"message": "Issue request rejected"}"
// @Failure 404 {object} object "{"error": "Request not found"}"
// @Failure 500 {object} object "{"error": "error message"}"
// @Router /requests/{id}/reject [post]
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
