package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online-library-system/database"
	"online-library-system/models"
)

func CreateRequestEvent(c *gin.Context) {
	var requestEvent models.RequestEvent
	if err := c.ShouldBindJSON(&requestEvent); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&requestEvent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, requestEvent)
}

func GetRequestEvents(c *gin.Context) {
	var requestEvents []models.RequestEvent
	database.DB.Find(&requestEvents)
	c.JSON(http.StatusOK, requestEvents)
}

func ApproveRequestEvent(c *gin.Context) {
	var requestEvent models.RequestEvent
	id := c.Param("id")
	if err := database.DB.First(&requestEvent, "req_id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Request event not found"})
		return
	}
	requestEvent.ApprovalDate = "today's date" // You can set it to the current date in real implementation
	if err := database.DB.Save(&requestEvent).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, requestEvent)
}
