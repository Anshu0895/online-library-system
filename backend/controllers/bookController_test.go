package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"online-library-system/models"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAddBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/books", AddBook)

	SetupTestDB()

	book := models.BookInventory{
		Title:           "Go Programming",
		ISBN:            "1234567890",
		Authors:         "John Doe",
		Publisher:       "Tech Books",
		TotalCopies:     5,
		AvailableCopies: 5,
	}

	body, _ := json.Marshal(book)
	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var response models.BookInventory
	json.Unmarshal(rr.Body.Bytes(), &response)

	if response.ISBN != book.ISBN {
		t.Errorf("handler returned unexpected ISBN: got %v want %v", response.ISBN, book.ISBN)
	}
}

func TestGetBooks(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/books", GetBooks)

	db := SetupTestDB()

	books := []models.BookInventory{
		{
			Title:           "Book One",
			ISBN:            "1111111111",
			Authors:         "Author One",
			Publisher:       "Publisher One",
			TotalCopies:     3,
			AvailableCopies: 3,
		},
		{
			Title:           "Book Two",
			ISBN:            "2222222222",
			Authors:         "Author Two",
			Publisher:       "Publisher Two",
			TotalCopies:     4,
			AvailableCopies: 4,
		},
	}

	db.Create(&books)

	req, _ := http.NewRequest("GET", "/books", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response []models.BookInventory
	json.Unmarshal(rr.Body.Bytes(), &response)

	if len(response) != len(books) {
		t.Errorf("handler returned wrong number of books: got %v want %v", len(response), len(books))
	}
}

func TestGetBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/books/:isbn", GetBook)

	db := SetupTestDB()

	book := models.BookInventory{
		Title:           "Test Book",
		ISBN:            "3333333333",
		Authors:         "Test Author",
		Publisher:       "Test Publisher",
		TotalCopies:     2,
		AvailableCopies: 2,
	}
	db.Create(&book)

	req, _ := http.NewRequest("GET", "/books/3333333333", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestRemoveBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.DELETE("/books/:isbn", RemoveBook)

	db := SetupTestDB()

	book := models.BookInventory{
		Title:           "Remove Book",
		ISBN:            "4444444444",
		Authors:         "Remove Author",
		Publisher:       "Remove Publisher",
		TotalCopies:     1,
		AvailableCopies: 1,
	}
	db.Create(&book)

	req, _ := http.NewRequest("DELETE", "/books/4444444444", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]string
	json.Unmarshal(rr.Body.Bytes(), &response)

	if response["message"] != "Available copy removed" {
		t.Errorf("handler returned unexpected message: got %v want %v", response["message"], "Available copy removed")
	}
}
func TestUpdateBook(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.PUT("/books/:isbn", UpdateBook)

	db := SetupTestDB()

	book := models.BookInventory{
		Title:           "Original Title",
		ISBN:            "5555555555",
		Authors:         "Original Author",
		Publisher:       "Original Publisher",
		TotalCopies:     3,
		AvailableCopies: 3,
	}
	db.Create(&book)

	updatedBook := models.BookInventory{
		Title:           "Updated Title",
		Authors:         "Updated Author",
		Publisher:       "Updated Publisher",
		TotalCopies:     5,
		AvailableCopies: 5,
	}

	body, _ := json.Marshal(updatedBook)
	req, _ := http.NewRequest("PUT", "/books/5555555555", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response models.BookInventory
	json.Unmarshal(rr.Body.Bytes(), &response)

	if response.Title != updatedBook.Title {
		t.Errorf("handler returned unexpected title: got %v want %v", response.Title, updatedBook.Title)
	}
	if response.Authors != updatedBook.Authors {
		t.Errorf("handler returned unexpected authors: got %v want %v", response.Authors, updatedBook.Authors)
	}
	if response.Publisher != updatedBook.Publisher {
		t.Errorf("handler returned unexpected publisher: got %v want %v", response.Publisher, updatedBook.Publisher)
	}
	if response.TotalCopies != updatedBook.TotalCopies {
		t.Errorf("handler returned unexpected total copies: got %v want %v", response.TotalCopies, updatedBook.TotalCopies)
	}
	if response.AvailableCopies != updatedBook.AvailableCopies {
		t.Errorf("handler returned unexpected available copies: got %v want %v", response.AvailableCopies, updatedBook.AvailableCopies)
	}
}
