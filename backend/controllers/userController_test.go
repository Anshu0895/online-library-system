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

func TestCreateAdmin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/create_admin", CreateAdmin)

	SetupTestDB()

	user := models.User{
		Name:          "Admin User",
		Email:         "admin@example.com",
		Password:      "adminpassword",
		ContactNumber: "9876543210",
	}

	body, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/create_admin", bytes.NewBuffer(body))
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

func TestGetUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/users/:id", GetUser)

	db := SetupTestDB()

	user := models.User{
		Name:          "Test User",
		Email:         "testuser@example.com",
		Password:      "testpassword",
		ContactNumber: "1234567890",
	}

	db.Create(&user)

	req, _ := http.NewRequest("GET", "/users/1", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response models.User
	json.Unmarshal(rr.Body.Bytes(), &response)

	if response.Email != user.Email {
		t.Errorf("handler returned unexpected email: got %v want %v", response.Email, user.Email)
	}
}

func TestGetUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/users", GetUsers)

	db := SetupTestDB()

	users := []models.User{
		{
			Name:          "User One",
			Email:         "userone@example.com",
			Password:      "password1",
			ContactNumber: "1111111111",
		},
		{
			Name:          "User Two",
			Email:         "usertwo@example.com",
			Password:      "password2",
			ContactNumber: "2222222222",
		},
	}

	db.Create(&users)

	req, _ := http.NewRequest("GET", "/users", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response []models.User
	json.Unmarshal(rr.Body.Bytes(), &response)

	if len(response) != len(users) {
		t.Errorf("handler returned wrong number of users: got %v want %v", len(response), len(users))
	}
}

func TestGetAdmins(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.GET("/admins", GetAdmins)

	db := SetupTestDB()

	admins := []models.User{
		{
			Name:          "Admin One",
			Email:         "adminone@example.com",
			Password:      "adminpassword1",
			ContactNumber: "3333333333",
			Role:          "LibraryAdmin",
		},
		{
			Name:          "Admin Two",
			Email:         "admintwo@example.com",
			Password:      "adminpassword2",
			ContactNumber: "4444444444",
			Role:          "LibraryAdmin",
		},
	}

	db.Create(&admins)

	req, _ := http.NewRequest("GET", "/admins", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response []models.User
	json.Unmarshal(rr.Body.Bytes(), &response)

	if len(response) != len(admins) {
		t.Errorf("handler returned wrong number of admins: got %v want %v", len(response), len(admins))
	}
}

func TestUpdateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.PUT("/users/:id", UpdateUser)

	db := SetupTestDB()

	user := models.User{
		Name:          "Test User",
		Email:         "testuser@example.com",
		Password:      "testpassword",
		ContactNumber: "1234567890",
	}

	db.Create(&user)

	updatedUser := models.User{
		Name:          "Updated User",
		Email:         "updateduser@example.com",
		ContactNumber: "9876543210",
	}

	body, _ := json.Marshal(updatedUser)
	req, _ := http.NewRequest("PUT", "/users/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response models.User
	json.Unmarshal(rr.Body.Bytes(), &response)

	if response.Email != updatedUser.Email {
		t.Errorf("handler returned unexpected email: got %v want %v", response.Email, updatedUser.Email)
	}
}

func TestDeleteUser(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.DELETE("/users/:id", DeleteUser)

	db := SetupTestDB()

	user := models.User{
		Name:          "Test User",
		Email:         "testuser@example.com",
		Password:      "testpassword",
		ContactNumber: "1234567890",
	}

	db.Create(&user)

	req, _ := http.NewRequest("DELETE", "/users/1", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var response map[string]string
	json.Unmarshal(rr.Body.Bytes(), &response)

	if response["message"] != "User deleted" {
		t.Errorf("handler returned unexpected message: got %v want %v", response["message"], "User deleted")
	}
}
