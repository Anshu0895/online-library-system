package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online-library-system/database"
	"online-library-system/models"
)

func CreateIssueRegistry(c *gin.Context) {
	var issueRegistry models.IssueRegistry
	if err := c.ShouldBindJSON(&issueRegistry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&issueRegistry).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, issueRegistry)
}

func GetIssueRegistries(c *gin.Context) {
	var issueRegistries []models.IssueRegistry
	database.DB.Find(&issueRegistries)
	c.JSON(http.StatusOK, issueRegistries)
}

func UpdateIssueStatus(c *gin.Context) {
	var issueRegistry models.IssueRegistry
	id := c.Param("id")
	if err := database.DB.First(&issueRegistry, "issue_id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Issue registry not found"})
		return
	}
	if err := c.ShouldBindJSON(&issueRegistry); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Save(&issueRegistry).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, issueRegistry)
}
