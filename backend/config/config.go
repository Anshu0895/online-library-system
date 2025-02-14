package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	//   DBConnectionString string
	ServerPort   string
	JWTSecretKey string
}

// LoadConfig loads the configuration settings from the .env file
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config := &Config{
		// DBConnectionString: os.Getenv("DB_CONNECTION_STRING"),
		ServerPort:   os.Getenv("SERVER_PORT"),
		JWTSecretKey: os.Getenv("JWT_SECRET_KEY"),
	}

	return config
}
