// package main

// import (
// 	"online-library-system/config"
// 	"online-library-system/database"
// 	"online-library-system/routes"

// 	"github.com/gin-gonic/gin"
// 	cors "github.com/rs/cors/wrapper/gin"
// )

// func main() {
// 	// Load the configuration settings
// 	cfg := config.LoadConfig()

// 	router := gin.Default()

// 	router.Use(cors.New(cors.Options{
// 		AllowedOrigins:   []string{cfg.AllowedOrigin}, // Frontend URL
// 		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
// 		AllowedHeaders:   []string{"Authorization", "Content-Type", "ApproverId"},
// 		ExposedHeaders:   []string{"Content-Length"},
// 		AllowCredentials: true,
// 	}))

// 	// Connect to Database
// 	database.Connect()

// 	// Setup Routes
// 	routes.AuthRoutes(router)
// 	routes.LibraryRoutes(router)
// 	routes.BookRoutes(router)
// 	routes.UserRoutes(router)
// 	routes.RequestRoutes(router)

// 	// Start Server
// 	router.Run(cfg.ServerPort)
// }
package main

import (
	// "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	cors "github.com/rs/cors/wrapper/gin"
	"log"
	"online-library-system/database"
	"online-library-system/routes"
	"os"
)

type Config struct {
	ServerPort   string
	JWTSecretKey string
}

var config *Config

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config = &Config{
		ServerPort:   os.Getenv("SERVER_PORT"),
		JWTSecretKey: os.Getenv("JWT_SECRET_KEY"),
	}
}

func main() {
	router := gin.Default()

	// CORS configuration
	// router.Use(cors.New(cors.Options{
	// 	AllowOrigins:     []string{"http://localhost:3000"},
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Content-Type", "Authorization", "ApproverID"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// }))

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"}, // Frontend URL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type", "ApproverId"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
	}))

	database.Connect()

	routes.AuthRoutes(router)
	routes.LibraryRoutes(router)
	routes.BookRoutes(router)
	routes.UserRoutes(router)
	routes.RequestRoutes(router)

	router.Run(":8080")
}
