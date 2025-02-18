package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online-library-system/database"
	"online-library-system/models"
)

func GetLibraries(c *gin.Context) {
	var libraries []models.Library
	if err := database.DB.Find(&libraries).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, libraries)
}
func CreateLibrary(c *gin.Context) {
	var library models.Library
	if err := c.ShouldBindJSON(&library); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check if the library already exists
	existingLibrary := models.Library{}
	if err := database.DB.Where("name = ?", library.Name).First(&existingLibrary).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Library with this name already exists"})
		return
	}

	// Check the user's role
	userID := c.GetUint("user_id")
	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve user"})
		return
	}

	if user.Role != "Owner" {
		c.JSON(http.StatusForbidden, gin.H{"error": "Only users with the role of 'Owner' can create a library"})
		return
	}
	if err := database.DB.Create(&library).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, library)
}
