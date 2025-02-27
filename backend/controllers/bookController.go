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

// @Summary Add a new book
// @Description Add a new book to the inventory
// @Tags books
// @Accept json
// @Security ApiKeyAuth
// @Produce json
// @Param book body models.BookInventory true "Book Data"
// @Success 201 {object} models.BookInventory
// @Failure 400 {object} object "{"error": "error message"}"
// @Failure 500 {object} object "{"error": "error message"}"
// @Router /books [post]
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

// @Summary Update an existing book
// @Description Update the details of an existing book by ISBN
// @Tags books
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param isbn path string true "Book ISBN"
// @Param book body models.BookInventory true "Updated Book Data"
// @Success 200 {object} models.BookInventory
// @Failure 400 {object} object "{"error": "error message"}"
// @Failure 404 {object} object "{"error": "Book not found"}"
// @Failure 500 {object} object "{"error": "error message"}"
// @Router /books/{isbn} [put]
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

// @Security ApiKeyAuth
// @Summary Get all books
// @Description Retrieve all books in the inventory
// @Tags books
// @Accept json
// @Produce json
// @Success 200 {array} models.BookInventory
// @Failure 500 {object} object "{"error": "error message"}"
// @Router /books [get]
func GetBooks(c *gin.Context) {
	var books []models.BookInventory
	if err := database.DB.Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, books)
}

// @Security ApiKeyAuth
// @Summary Get a book by ISBN
// @Description Retrieve a book by its ISBN
// @Tags books
// @Accept json
// @Produce json
// @Param isbn path string true "Book ISBN"
// @Success 200 {object} models.BookInventory
// @Failure 404 {object} object "{"error": "Book not found"}"
// @Failure 500 {object} object "{"error": "error message"}"
// @Router /books/{isbn} [get]
func GetBook(c *gin.Context) {
	var book models.BookInventory
	isbn := c.Param("isbn")
	if err := database.DB.First(&book, "isbn = ?", isbn).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
		return
	}
}

// @Security ApiKeyAuth
// @Summary Remove a book
// @Description Remove a book from the inventory by ISBN
// @Tags books
// @Accept json
// @Produce json
// @Param isbn path string true "Book ISBN"
// @Success 200 {object} object "{"message": "Available copy removed"}"
// @Failure 400 {object} object "{"error": "No available copies to remove"}"
// @Failure 404 {object} object "{"error": "Book not found"}"
// @Failure 500 {object} object "{"error": "error message"}"
// @Router /books/{isbn} [delete]
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

// @Security ApiKeyAuth
// @Summary Search for books
// @Description Search for books by title, author, publisher, or status
// @Tags books
// @Accept json
// @Produce json
// @Param title query string false "Book Title"
// @Param author query string false "Book Author"
// @Param publisher query string false "Book Publisher"
// @Param status query string false "Book Status"
// @Success 200 {array} models.BookInventory
// @Failure 500 {object} object "{"error": "error message"}"
// @Router /books/search [get]

func SearchBooks(c *gin.Context) {
	title := c.Query("title")
	author := c.Query("author")
	publisher := c.Query("publisher")
	status := c.Query("status")

	var books []models.BookInventory

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
