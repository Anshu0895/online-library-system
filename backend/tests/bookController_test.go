package tests

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"online-library-system/controllers"
	"strings"
	"testing"
)

func TestAddBook(t *testing.T) {
	router := gin.Default()
	router.POST("/books", controllers.AddBook)

	bookJSON := `{"ISBN": "1234567890", "LibID": 1, "Title": "Test Book", "Authors": "Test Author", "Publisher": "Test Publisher", "Version": "1.0", "TotalCopies": 5, "AvailableCopies": 5}`
	req, _ := http.NewRequest("POST", "/books", strings.NewReader(bookJSON))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
}
