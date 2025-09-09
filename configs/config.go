package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Smtp SmtpConfig
	Port string
}

type SmtpConfig struct {
	Email    string
	Password string
	Server   string
}

func NewConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return &Config{
		Smtp: SmtpConfig{
			Email:    os.Getenv("SMTP_LOGIN"),
			Password: os.Getenv("SMTP_PASSWORD"),
			Server:   os.Getenv("SMTP_SERVER"),
		},

		Port: os.Getenv("SERVER_PORT"),
	}
}
