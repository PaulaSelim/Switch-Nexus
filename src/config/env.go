package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SSHHost     string
	SSHUser     string
	SSHPassword string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	config := &Config{
		SSHHost:     os.Getenv("SSH_HOST"),
		SSHUser:     os.Getenv("SSH_USER"),
		SSHPassword: os.Getenv("SSH_PASSWORD"),
	}

	if config.SSHHost == "" || config.SSHUser == "" || config.SSHPassword == "" {
		log.Fatalf("Required environment variables are not set")
	}

	return config
}
