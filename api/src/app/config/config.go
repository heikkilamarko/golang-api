// Package config provides application configuration functionality.
package config

import (
	"fmt"
	"os"
)

// Config struct
type Config struct {
	Port        string
	DBHost      string
	DBPort      string
	DBName      string
	DBUsername  string
	DBPassword  string
	DBSSLMode   string
	APIKey      string
	CORSEnabled bool
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

// GetValue func
func GetValue(key, fallback string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return fallback
	}
	return value
}
