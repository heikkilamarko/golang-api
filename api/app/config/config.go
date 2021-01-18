// Package config provides application configuration functionality.
package config

import (
	"fmt"
	"os"
	"time"
)

// Config struct
type Config struct {
	Port           string
	DBHost         string
	DBPort         string
	DBName         string
	DBUsername     string
	DBPassword     string
	DBSSLMode      string
	APIKey         string
	CORSEnabled    bool
	RequestTimeout time.Duration
}

// New func
func New() *Config {
	return &Config{}
}

// Load method
func (c *Config) Load() {
	c.Port = getEnv("APP_PORT", "8080")
	c.DBHost = getEnv("APP_DB_HOST", "")
	c.DBName = getEnv("APP_DB_NAME", "")
	c.DBPort = getEnv("APP_DB_PORT", "5432")
	c.DBUsername = getEnv("APP_DB_USERNAME", "")
	c.DBPassword = getEnv("APP_DB_PASSWORD", "")
	c.DBSSLMode = getEnv("APP_DB_SSLMODE", "require")
	c.APIKey = getEnv("APP_API_KEY", "")
	c.CORSEnabled = getEnv("APP_CORS_ENABLED", "") == "true"

	var err error
	if c.RequestTimeout, err = time.ParseDuration(getEnv("APP_REQUEST_TIMEOUT", "10s")); err != nil {
		c.RequestTimeout = 10 * time.Second
	}
}

// ServerAddr method
func (c *Config) ServerAddr() string {
	return fmt.Sprintf(":%s", c.Port)
}

// PostgresConnectionString method
func (c *Config) PostgresConnectionString() string {
	return fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=%s",
		c.DBHost,
		c.DBPort,
		c.DBName,
		c.DBUsername,
		c.DBPassword,
		c.DBSSLMode)
}

func getEnv(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}
