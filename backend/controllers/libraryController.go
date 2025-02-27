package controllers

import (
	"net/http"
	"online-library-system/database"
	"online-library-system/models"

	"github.com/gin-gonic/gin"
)

// @Summary Get all libraries
// @Description Retrieve all libraries
// @Tags libraries
// @Accept json
// @Produce json
// @Security ApiKeyAuth

// @Success 200 {array} models.Library
// @Failure 500 {object} object "{"error": "error message"}"
// @Router /libraries [get]
func GetLibraries(c *gin.Context) {
	var libraries []models.Library
	if err := database.DB.Find(&libraries).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, libraries)
}

// @Summary Create a new library
// @Description Create a new library (only users with the role of 'Owner' can create a library)
// @Tags libraries
// @Security ApiKeyAuth
// @Accept json
// @Produce json
// @Param library body models.Library true "Library Data"
// @Success 201 {object} models.Library
// @Failure 400 {object} object "{"error": "error message"}"
// @Failure 409 {object} object "{"error": "Library with this name already exists"}"
// @Failure 403 {object} object "{"error": "Only users with the role of 'Owner' can create a library"}"
// @Failure 500 {object} object "{"error": "error message"}"
// @Router /libraries [post]
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

// @Summary Delete a library
// @Description Delete a library by ID
// @Tags libraries
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Library ID"
// @Success 200 {object} object "{"message": "library deleted successfuly"}"
// @Failure 404 {object} object "{"message": "library not found"}"
// @Failure 500 {object} object "{"message": "error message"}"
// @Router /libraries/{id} [delete]
func DeleteLibrary(c *gin.Context) {
	// Get library ID from the URL parameter and convert it to uint
	var library models.Library
	id := c.Param("id")

	if err := database.DB.First(&library, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "library not found"})
		return
	}
	if err := database.DB.Delete(&library).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "library deleted successfuly"})

}
