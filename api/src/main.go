package main

import (
	"os"
	"products-api/app"
	"products-api/app/config"

	"github.com/rs/zerolog"
)

func main() {

	c := &config.Config{
		Port:        config.GetValue("APP_PORT", "8080"),
		DBHost:      config.GetValue("APP_DB_HOST", ""),
		DBName:      config.GetValue("APP_DB_NAME", ""),
		DBPort:      config.GetValue("APP_DB_PORT", "5432"),
		DBUsername:  config.GetValue("APP_DB_USERNAME", ""),
		DBPassword:  config.GetValue("APP_DB_PASSWORD", ""),
		DBSSLMode:   config.GetValue("APP_DB_SSLMODE", "require"),
		APIKey:      config.GetValue("APP_API_KEY", ""),
		CORSEnabled: config.GetValue("APP_CORS_ENABLED", "") == "true",
	}

	l := zerolog.New(os.Stderr).With().Timestamp().Logger()

	a := app.New(c, &l)

	a.Run()
}
