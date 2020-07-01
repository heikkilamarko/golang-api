package config

import "os"

// Config global var
var Config *AppConfig = &AppConfig{}

// AppConfig struct
type AppConfig struct {
	Port        string
	DBHost      string
	DBPort      string
	DBName      string
	DBUsername  string
	DBPassword  string
	APIKey      string
	CorsEnabled bool
}

// Load method
func Load() {
	Config.Port = os.Getenv("APP_PORT")
	Config.DBHost = os.Getenv("APP_DB_HOST")
	Config.DBName = os.Getenv("APP_DB_NAME")
	Config.DBPort = os.Getenv("APP_DB_PORT")
	Config.DBUsername = os.Getenv("APP_DB_USERNAME")
	Config.DBPassword = os.Getenv("APP_DB_PASSWORD")
	Config.APIKey = os.Getenv("APP_API_KEY")
	Config.CorsEnabled = os.Getenv("APP_CORS_ENABLED") == "true"
}
