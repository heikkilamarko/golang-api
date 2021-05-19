// Package config provides application configuration functionality.
package config

import "products-api/app/utils"

// Config struct
type Config struct {
	App                string `json:"app"`
	Address            string `json:"address"`
	APIKey             string `json:"api_key"`
	APIKeyHeader       string `json:"api_key_header"`
	DBConnectionString string `json:"db_connection_string"`
	CORSEnabled        bool   `json:"cors_enabled"`
	LogLevel           string `json:"log_level"`
}

// Load func
func Load() *Config {
	return &Config{
		App:                utils.Env("APP_NAME", ""),
		Address:            utils.Env("APP_ADDRESS", ":8080"),
		APIKey:             utils.Env("APP_API_KEY", ""),
		APIKeyHeader:       utils.Env("APP_API_KEY_HEADER", "X-Api-Key"),
		DBConnectionString: utils.Env("APP_DB_CONNECTION_STRING", ""),
		CORSEnabled:        utils.Env("APP_CORS_ENABLED", "") == "true",
		LogLevel:           utils.Env("APP_LOG_LEVEL", "warn"),
	}
}
