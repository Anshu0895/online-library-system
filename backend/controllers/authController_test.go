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
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect to the database")
	}
	db.AutoMigrate(&models.Library{}, &models.User{}, &models.BookInventory{}, &models.RequestEvent{}, &models.IssueRegistry{})
	database.DB = db
	return db
}

func TestSignup(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/signup", Signup)

	SetupTestDB()

	user := models.User{
		Name:          "John Doe",
		Email:         "johndoe@example.com",
		Password:      "password123",
		ContactNumber: "1234567890",
		Role:          "Admin",
	}

	body, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	var response models.User
	json.Unmarshal(rr.Body.Bytes(), &response)

	if response.Email != user.Email {
		t.Errorf("handler returned unexpected email: got %v want %v", response.Email, user.Email)
	}
}

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/login", Login)

	SetupTestDB()

	// Create a test user
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	testUser := models.User{
		Name:          "John Doe",
		Email:         "johndoe@example.com",
		Password:      string(hashedPassword),
		ContactNumber: "1234567890",
	}
	database.DB.Create(&testUser)

	credentials := map[string]string{
		"email":    "johndoe@example.com",
		"password": "password123",
	}

	body, _ := json.Marshal(credentials)
	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]interface{}
	json.Unmarshal(rr.Body.Bytes(), &response)

	if response["token"] == "" {
		t.Errorf("handler did not return a token")
	}
}
