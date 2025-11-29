package configs

import "os"

type Config struct {
	Email    string
	Password string
	Address  string
	Port     string
}

func LoadConfig() *Config {
	return &Config{
		Email:    os.Getenv("EMAIL"),
		Password: os.Getenv("PASSWORD"),
		Address:  os.Getenv("ADDRESS"),
		Port:     os.Getenv("PORT"),
	}
}
