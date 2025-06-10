package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBConnectionString string
	AppPort            string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("WARNING: Could not load .env file. Using environment variables directly.")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("DB_SSLMODE")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, sslMode)

	appPort := os.Getenv("APP_PORT")
	if appPort == "" {
		appPort = "8080" // fallback
	}

	return &Config{
		DBConnectionString: connStr,
		AppPort:            appPort,
	}
}
