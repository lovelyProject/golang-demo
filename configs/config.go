package configs

import (
	"github.com/joho/godotenv"
	"os"
)

type SmptConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	From     string
}

type DbConfig struct {
	Dsn string
}
type Config struct {
	Port string
	Db   DbConfig
}

func NewConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}

	return &Config{
		Port: os.Getenv("SERVER_PORT"),
		Db: DbConfig{
			Dsn: os.Getenv("DB_DSN"),
		},
	}
}
