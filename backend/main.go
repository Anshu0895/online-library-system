package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	cors "github.com/rs/cors/wrapper/gin"
	"log"
	"online-library-system/database"
	// "online-library-system/docs"
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

	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")
}
