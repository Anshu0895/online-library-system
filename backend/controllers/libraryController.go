package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online-library-system/database"
	"online-library-system/models"
)

func CreateLibrary(c *gin.Context) {
	var library models.Library
	if err := c.ShouldBindJSON(&library); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&library).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, library)
}
