package main

import (
	"os"
	"products-api/app"
	"products-api/app/config"

	"github.com/rs/zerolog"
)

func main() {

	c := &config.Config{
		Port:        os.Getenv("APP_PORT"),
		DBHost:      os.Getenv("APP_DB_HOST"),
		DBName:      os.Getenv("APP_DB_NAME"),
		DBPort:      os.Getenv("APP_DB_PORT"),
		DBUsername:  os.Getenv("APP_DB_USERNAME"),
		DBPassword:  os.Getenv("APP_DB_PASSWORD"),
		DBSSLMode:   os.Getenv("APP_DB_SSLMODE"),
		APIKey:      os.Getenv("APP_API_KEY"),
		CorsEnabled: os.Getenv("APP_CORS_ENABLED") == "true",
	}

	l := zerolog.New(os.Stderr).With().Timestamp().Logger()

	a := app.New(c, &l)

	a.Run()
}
