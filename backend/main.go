package main

import (
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
	"online-library-system/config"
	"online-library-system/database"
	"online-library-system/routes"
)

func main() {
	// Load the configuration settings
	cfg := config.LoadConfig()

	router := gin.Default()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{cfg.AllowedOrigin}, // Frontend URL
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Enable CORS
	// router.Use(cors.Default())

	// Connect to Database
	database.Connect()

	// Setup Routes
	routes.AuthRoutes(router)
	routes.LibraryRoutes(router)
	routes.BookRoutes(router)
	routes.UserRoutes(router)
	routes.RequestRoutes(router)

	// Start Server
	router.Run(cfg.ServerPort)
}
