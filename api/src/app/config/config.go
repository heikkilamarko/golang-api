// Package config provides application configuration functionality.
package config

import "fmt"

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
	CorsEnabled bool
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
