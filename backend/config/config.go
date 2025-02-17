package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	ServerPort    string
	JWTSecretKey  string
	AllowedOrigin string
}

// LoadConfig loads the configuration settings from the .env file
func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config := &Config{
		ServerPort:    os.Getenv("SERVER_PORT"),
		JWTSecretKey:  os.Getenv("JWT_SECRET_KEY"),
		AllowedOrigin: os.Getenv("ALLOWED_ORIGIN"),
	}

	return config
}
