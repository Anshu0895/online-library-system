package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"online-library-system/database"
	"online-library-system/models"
	"os"
	"regexp"
	"strconv"
	"time"
)

type Config struct {
	//   DBConnectionString string
	ServerPort   string
	JWTSecretKey string
}

var config *Config
var jwtKey []byte

func init() {
	err := godotenv.Load()
	if err != nil {
	}

	config = &Config{
		// DBConnectionString: os.Getenv("DB_CONNECTION_STRING"),
		ServerPort:   os.Getenv("SERVER_PORT"),
		JWTSecretKey: os.Getenv("JWT_SECRET_KEY"),
	}

	jwtKey = []byte(config.JWTSecretKey)
}

type Claims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.StandardClaims
}

// @Summary Sign up a new user
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "User Data"
// @Success 201 {object} models.User
// @Failure 400 {object} gin.H{"error": "error message"}
// @Failure 500 {object} gin.H{"error": "error message"}
// @Router /signup [post]

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
	existingUser, err := GetUserByEmail(user.Email)
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

	if err := CreateUser(&user); err != nil {
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

// @Summary Log in a user
// @Description Log in a user and return a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param credentials body models.Credentials true "User Credentials"
// @Success 200 {object} gin.H{"token": "token string", "role": "user role", "id": "user ID"}
// @Failure 400 {object} gin.H{"error": "error message"}
// @Failure 401 {object} gin.H{"error": "error message"}
// @Failure 500 {object} gin.H{"error": "error message"}

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

	user, err := GetUserByEmail(credentials.Email)
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
	c.JSON(http.StatusOK, gin.H{"token": tokenString, "role": user.Role, "id": user.ID})
}

// Service function to create a user in the database
func CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

// Service function to get a user by email from the database
func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("email = ?", email).First(&user)
	return &user, result.Error
}
