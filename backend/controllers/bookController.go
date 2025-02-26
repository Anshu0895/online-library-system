package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	// "gorm.io/gorm"
	// "gorm.io/gorm/logger"
	"net/http"
	"online-library-system/database"
	"online-library-system/models"
	"strings"
)

func AddBook(c *gin.Context) {
	var book models.BookInventory
	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	existingBook := models.BookInventory{}
	if err := database.DB.Where("isbn = ?", book.ISBN).First(&existingBook).Error; err == nil {
		existingBook.TotalCopies += book.TotalCopies
		existingBook.AvailableCopies += book.AvailableCopies
		if err := database.DB.Save(&existingBook).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, existingBook)
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
func GetBooks(c *gin.Context) {
	var books []models.BookInventory
	if err := database.DB.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	var book models.BookInventory
	isbn := c.Param("isbn")
	if err := database.DB.First(&book, "isbn = ?", isbn).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
}
func RemoveBook(c *gin.Context) {
	var book models.BookInventory
	isbn := c.Param("isbn")

	// Check if the book exists
	if err := database.DB.First(&book, "isbn = ?", isbn).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}

	// Check if there are available copies to decrement
	if book.AvailableCopies > 0 {
		book.AvailableCopies--
		book.TotalCopies--
		if err := database.DB.Save(&book).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"message": "Available copy removed"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No available copies to remove"})
	}
}

func SearchBooks(c *gin.Context) {
	title := c.Query("title")
	author := c.Query("author")
	publisher := c.Query("publisher")
	status := c.Query("status")

	var books []models.BookInventory
	// query := database.DB.Session(&gorm.Session{
	// 	Logger: logger.Default.LogMode(logger.Info),
	// })
	query := database.DB
	if title != "" {
		query = query.Where("LOWER(title) LIKE ?", "%"+strings.ToLower(title)+"%")
	}
	if author != "" {
		query = query.Where("LOWER(authors) LIKE ?", "%"+strings.ToLower(author)+"%")
	}
	if publisher != "" {
		query = query.Where("LOWER(publisher) LIKE ?", "%"+strings.ToLower(publisher)+"%")
	}
	if status != "" {
		if status == "available" {
			query = query.Where("available_copies > ?", 0)

		} else {
			query = query.Where("available_copies = ?", 0)

		}
	}

	if err := query.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Debug log
	fmt.Println("Books found:", books)

	c.JSON(http.StatusOK, books)
}
