package controllers

import (
	"net/http"
	"online-library-system/database"
	"online-library-system/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// @Summary Create a new admin
// @Description Create a new admin user with role "LibraryAdmin"
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param user body models.User true "User Data"
// @Success 201 {object} models.User
// @Failure 400 {object} object "{"error": "error message"}"
// @Failure 500 {object} object "{"error": "error message"}"
// @Router /users/admin [post]
func CreateAdmin(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Hash the password before storing it
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword) // Store the hashed password

	user.Role = "LibraryAdmin"

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Remove the hashed password from the response for security
	user.Password = ""

	c.JSON(http.StatusCreated, user)
}

// @Summary Get a user by ID
// @Description Retrieve a user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} models.User
// @Failure 404 {object} object "{"error": "User not found"}"
// @Failure 500 {object} object "{"error": "error message"}"
// @Router /users/{id} [get]
func GetUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := database.DB.First(&user, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Summary Get all users
// @Description Retrieve all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} object "{"error": "error message"}"
// @Router /users [get]
func GetUsers(c *gin.Context) {
	var users []models.User
	if err := database.DB.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

// @Summary Get all admins
// @Description Retrieve all admin users with role "LibraryAdmin"
// @Tags users
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} models.User
// @Failure 500 {object} object "{"error": "error message"}"
// @Router /users/admins [get]
func GetAdmins(c *gin.Context) {
	var admins []models.User

	if err := database.DB.Where("role = ?", "LibraryAdmin").Find(&admins).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, admins)
}

// @Summary Update a user
// @Description Update a user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body models.User true "Updated User Data"
// @Success 200 {object} models.User
// @Failure 400 {object} object "{"error": "error message"}"
// @Failure 404 {object} object "{"error": "User not found"}"
// @Failure 500 {object} object "{"error": "error message"}"
// @Router /users/{id} [put]
func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := database.DB.First(&user, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Summary Delete a user
// @Description Delete a user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} object "{"message": "User deleted"}"
// @Failure 404 {object} object "{"error": "User not found"}"
// @Failure 500 {object} object "{"error": "error message"}"
// @Router /users/{id} [delete]
func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	if err := database.DB.First(&user, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	if err := database.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
