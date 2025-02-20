package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
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
	fmt.Println("Received book_id:", request.BookID) // Debug log
	// Check availability of the book
	var book models.BookInventory
	if err := database.DB.First(&book, "isbn = ?", request.BookID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	if book.AvailableCopies <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book not available"})
		return
	}

	// Create Issue Request
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
//Get Pending Requests
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

	// Approve Request
	request.ApprovalDate = time.Now()
	request.ApproverID = 3
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
