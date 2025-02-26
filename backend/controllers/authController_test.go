
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

	tests := []struct {
		name     string
		user     models.User
		wantCode int
		wantBody string
	}{
		{
			name: "Valid user",
			user: models.User{
				Name:          "John Doe",
				Email:         "johndoe@example.com",
				Password:      "password123",
				ContactNumber: "1234567890",
				Role:          "Admin",
			},
			wantCode: http.StatusCreated,
		},
		{
			name: "Invalid email",
			user: models.User{
				Name:          "John Doe",
				Email:         "invalid-email",
				Password:      "password123",
				ContactNumber: "1234567890",
				Role:          "Admin",
			},
			wantCode: http.StatusBadRequest,
			wantBody: "Invalid email address",
		},
		{
			name: "Short password",
			user: models.User{
				Name:          "John Doe",
				Email:         "johndoe@example.com",
				Password:      "short",
				ContactNumber: "1234567890",
				Role:          "Admin",
			},
			wantCode: http.StatusBadRequest,
			wantBody: "Password must be at least 8 characters",
		},
		{
			name: "Invalid contact number",
			user: models.User{
				Name:          "John Doe",
				Email:         "johndoe@example.com",
				Password:      "password123",
				ContactNumber: "invalid-number",
				Role:          "Admin",
			},
			wantCode: http.StatusBadRequest,
			wantBody: "Contact number must be 10 digits",
		},
		{
			name: "User already exists",
			user: models.User{
				Name:          "John Doe",
				Email:         "johndoe@example.com",
				Password:      "password123",
				ContactNumber: "1234567890",
				Role:          "Admin",
			},
			wantCode: http.StatusBadRequest,
			wantBody: "User with this email already exists",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create the user if the test case is "User already exists"
			if tt.name == "User already exists" {
				existingUser := tt.user
				hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(existingUser.Password), bcrypt.DefaultCost)
				existingUser.Password = string(hashedPassword)
				database.DB.Create(&existingUser)
			}

			body, _ := json.Marshal(tt.user)
			req, _ := http.NewRequest("POST", "/signup", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.wantCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.wantCode)
			}

			if tt.wantBody != "" && !jsonContains(rr.Body.Bytes(), tt.wantBody) {
				t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), tt.wantBody)
			}
		})
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
		Role:          "Admin",
	}
	database.DB.Create(&testUser)

	tests := []struct {
		name     string
		creds    map[string]string
		wantCode int
		wantBody string
	}{
		{
			name: "Valid login",
			creds: map[string]string{
				"email":    "johndoe@example.com",
				"password": "password123",
			},
			wantCode: http.StatusOK,
		},
		{
			name: "Invalid email",
			creds: map[string]string{
				"email":    "invalid@example.com",
				"password": "password123",
			},
			wantCode: http.StatusUnauthorized,
			wantBody: "Invalid email or password",
		},
		{
			name: "Incorrect password",
			creds: map[string]string{
				"email":    "johndoe@example.com",
				"password": "wrongpassword",
			},
			wantCode: http.StatusUnauthorized,
			wantBody: "Invalid email or password",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			body, _ := json.Marshal(tt.creds)
			req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			rr := httptest.NewRecorder()
			router.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.wantCode {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.wantCode)
			}

			if tt.wantBody != "" && !jsonContains(rr.Body.Bytes(), tt.wantBody) {
				t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), tt.wantBody)
			}
		})
	}
}

// Helper function to check if JSON response contains a substring
func jsonContains(body []byte, substring string) bool {
	var result map[string]interface{}
	json.Unmarshal(body, &result)
	for _, v := range result {
		if str, ok := v.(string); ok && str == substring {
			return true
		}
	}
	return false
}
