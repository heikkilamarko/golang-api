// Package config provides application configuration functionality.
package config

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

// Config struct
type Config struct {
	Address            string
	APIKey             string
	APIKeyHeader       string
	DBConnectionString string
	CORSEnabled        bool
	LogLevel           zerolog.Level
	RequestTimeout     time.Duration
}

// New func
func New() *Config {
	return &Config{}
}

// Load method
func (c *Config) Load() {
	c.Address = getEnv("APP_ADDRESS", ":8080")
	c.APIKey = getEnv("APP_API_KEY", "")
	c.APIKeyHeader = getEnv("APP_API_KEY_HEADER", "X-Api-Key")
	c.DBConnectionString = getEnv("APP_DB_CONNECTION_STRING", "")
	c.CORSEnabled = getEnv("APP_CORS_ENABLED", "") == "true"

	var err error

	if c.LogLevel, err = zerolog.ParseLevel(getEnv("APP_LOG_LEVEL", "warn")); err != nil {
		c.LogLevel = zerolog.WarnLevel
	}

	if c.RequestTimeout, err = time.ParseDuration(getEnv("APP_REQUEST_TIMEOUT", "10s")); err != nil {
		c.RequestTimeout = 10 * time.Second
	}
}

func getEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}
