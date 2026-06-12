package config

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// Config holds the application configuration parameters.
// The fields are populated from environment variables at startup.
type Config struct {
	Port        string
	ContentPath string
	JWTSecret   string
}

// Load reads configuration from environment variables, applies default values,
// validates required fields, and returns a usable *Config.
//
// Required:
//   - JWT_SECRET
//
// Optional (with defaults):
//   - PORT (default: 8080)
//   - CONTENT_PATH (default: ./content)
func Load() (*Config, error) {
	cfg := &Config{
		Port:        getEnv("PORT", "8080"),
		ContentPath: getEnv("CONTENT_PATH", "./content"),
		JWTSecret:   getEnv("JWT_SECRET", ""),
	}

	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return cfg, nil
}

// validate ensures all required fields are populated
// and contain sane values.
func (c *Config) validate() error {
	if strings.TrimSpace(c.JWTSecret) == "" {
		return errors.New("missing required environment variable: JWT_SECRET")
	}

	if strings.TrimSpace(c.Port) == "" {
		return errors.New("environment variable PORT cannot be empty or whitespace")
	}

	if strings.TrimSpace(c.ContentPath) == "" {
		return errors.New("environment variable CONTENT_PATH cannot be empty or whitespace")
	}

	return nil
}

// getEnv fetches an environment variable by key.
// If the variable is unset or empty, it returns the fallback value.
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}