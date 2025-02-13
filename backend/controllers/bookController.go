package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"online-library-system/database"
	"online-library-system/models"
)

func AddBook(c *gin.Context) {
	var book models.BookInventory
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, book)
}

func UpdateBook(c *gin.Context) {
	var book models.BookInventory
	isbn := c.Param("isbn")
	if err := database.DB.First(&book, "isbn = ?", isbn).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, book)
}

func RemoveBook(c *gin.Context) {
	var book models.BookInventory
	isbn := c.Param("isbn")
	if err := database.DB.First(&book, "isbn = ?", isbn).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
	if err := database.DB.Delete(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Book removed"})
}
