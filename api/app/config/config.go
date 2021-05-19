// Package config provides application configuration functionality.
package config

import (
	"encoding/json"
	"os"
	"time"

	"github.com/rs/zerolog"
)

// Config struct
type Config struct {
	App                string        `json:"app"`
	Address            string        `json:"address"`
	APIKey             string        `json:"api_key"`
	APIKeyHeader       string        `json:"api_key_header"`
	DBConnectionString string        `json:"db_connection_string"`
	CORSEnabled        bool          `json:"cors_enabled"`
	LogLevel           zerolog.Level `json:"log_level"`
	RequestTimeout     time.Duration `json:"request_timeout"`
}

// New func
func New() *Config {
	return &Config{}
}

const redactedString = "[redacted]"

// String method
func (c *Config) String() string {
	cc := *c
	cc.APIKey = redactedString
	cc.DBConnectionString = redactedString
	if b, err := json.Marshal(cc); err == nil {
		return string(b)
	}
	return ""
}

// Load method
func (c *Config) Load() {
	c.App = getEnv("APP_NAME", "")
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
