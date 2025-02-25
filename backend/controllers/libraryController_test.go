package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"online-library-system/database"
	"online-library-system/models"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Helper function to create a test router
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/libraries", GetLibraries)
	r.POST("/libraries", CreateLibrary)
	r.DELETE("/libraries/:id", DeleteLibrary)
	return r
}

func TestGetLibraries(t *testing.T) {
	SetupTestDB()
	router := SetupRouter()

	// Insert test data
	database.DB.Create(&models.Library{Name: "Test Library"})

	// Perform GET request
	req, _ := http.NewRequest("GET", "/libraries", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Test Library")
}

func TestCreateLibrary(t *testing.T) {
	SetupTestDB() // Ensure test DB is initialized

	// Create an owner user
	owner := models.User{ID: 1, Role: "Owner"}
	database.DB.Create(&owner)

	// Create a library request
	libraryData := map[string]string{"name": "New Library"}
	jsonValue, _ := json.Marshal(libraryData)

	// Create a new request
	req, _ := http.NewRequest("POST", "/libraries", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	// Create a test Gin context
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// Simulate user authentication
	c.Request = req
	c.Set("user_id", uint(1)) // Ensure user_id is set

	// Call the handler manually with the updated context
	CreateLibrary(c)

	// Assertions
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "New Library")
}

func TestCreateDuplicateLibrary(t *testing.T) {
	SetupTestDB()
	router := SetupRouter()

	// Insert existing library
	database.DB.Create(&models.Library{Name: "Duplicate Library"})

	// Create a library request with the same name
	libraryData := map[string]string{"name": "Duplicate Library"}
	jsonValue, _ := json.Marshal(libraryData)

	req, _ := http.NewRequest("POST", "/libraries", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusConflict, w.Code)
	assert.Contains(t, w.Body.String(), "Library with this name already exists")
}

func TestDeleteLibrary(t *testing.T) {
	SetupTestDB()
	router := SetupRouter()

	// Insert test library
	library := models.Library{Name: "Library to Delete"}
	database.DB.Create(&library)

	// Perform DELETE request
	req, _ := http.NewRequest("DELETE", "/libraries/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "library deleted successfuly")
}

func TestDeleteNonExistentLibrary(t *testing.T) {
	SetupTestDB()
	router := SetupRouter()

	// Perform DELETE request for a non-existing library
	req, _ := http.NewRequest("DELETE", "/libraries/999", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assertions
	assert.Equal(t, http.StatusNotFound, w.Code)
	assert.Contains(t, w.Body.String(), "library not found")
}
