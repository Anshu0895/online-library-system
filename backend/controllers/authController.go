package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"online-library-system/config"
	"online-library-system/models"
	"online-library-system/services"
	"regexp"
	"strconv"
	"time"
)

var cfg = config.LoadConfig()
var jwtKey = []byte(cfg.JWTSecretKey)

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

func Signup(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Debug: Print received user data
	fmt.Printf("Received user data: %+v\n", user)
	// Validate email
	if !isValidEmail(user.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid email address"})
		return
	}

	// Validate password
	if len(user.Password) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 8 characters"})
		return
	}

	// Validate contact number
	if len(user.ContactNumber) != 10 || !isNumeric(user.ContactNumber) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Contact number must be 10 digits"})
		return
	}
	// Debug: Print contact number validation
	fmt.Printf("Contact Number Validation: %s\n", user.ContactNumber)
	// Check if user already exists
	existingUser, err := services.GetUserByEmail(user.Email)
	if err == nil && existingUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with this email already exists"})
		return
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}
	user.Password = string(hashedPassword)

	if err := services.CreateUser(&user); err != nil {
		fmt.Printf("Failed to create user: %v\n", err) //debug
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("User created successfully") //debug
	c.JSON(http.StatusCreated, user)
}

// Helper function to validate email
func isValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

// Helper function to check if a string is numeric
func isNumeric(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}
func Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//debug
	fmt.Printf("Received credentials: %+v\n", credentials)

	user, err := services.GetUserByEmail(credentials.Email)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)) != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}
	//debug
	fmt.Printf("Retrieved user: %+v\n", user)
	// Generate JWT token
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: user.ID,
		Role:   user.Role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}
	// Debug: Print generated token
	fmt.Printf("Generated token: %s\n", tokenString)
	c.JSON(http.StatusOK, gin.H{"token": tokenString, "role": user.Role})
}
