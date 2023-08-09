package util

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	SecretKey string
	Port      string
}

func LoadConfig() Config {
	return Config{
		SecretKey: os.Getenv("SECRET_KEY"),
		Port:      os.Getenv("PORT"),
	}
}
