package main

import (
	"github.com/gin-gonic/gin"
	"online-library-system/config"
	"online-library-system/database"
	"online-library-system/routes"
)

func main() {
	// Load the configuration settings
	cfg := config.LoadConfig()

	router := gin.Default()

	// Connect to Database
	database.Connect()

	// Setup Routes
	routes.AuthRoutes(router)
	routes.LibraryRoutes(router)
	routes.BookRoutes(router)
	routes.UserRoutes(router)
	routes.RequestRoutes(router)
	routes.IssueRoutes(router)

	// Start Server
	router.Run(cfg.ServerPort)
}
