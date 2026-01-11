package configs

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Password string
	Address  string
	Port     string
	Email    string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := &Config{
		Password: os.Getenv("PASSWORD"),
		Address:  os.Getenv("ADDRESS"),
		Port:     os.Getenv("PORT"),
		Email:    os.Getenv("EMAIL"),
	}

	if cfg.Password == "" {
		return nil, fmt.Errorf("environment variable PASSWORD is not set")
	}
	if cfg.Address == "" {
		return nil, fmt.Errorf("environment variable ADDRESS is not set")
	}
	if cfg.Port == "" {
		return nil, fmt.Errorf("environment variable PORT is not set")
	}
	if cfg.Email == "" {
		return nil, fmt.Errorf("environment variable EMAIL is not set")
	}

	return cfg, nil
}
