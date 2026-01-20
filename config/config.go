package config

import "os"

type AppConfig struct {
	DatabaseURL string
	Port        string

	// JWTSecret      string
	// AdminEmail     string
	// Environment    string
	// MaxRequestSize int64
}

// exported - uppercase
func Load() *AppConfig {
	return &AppConfig{
		DatabaseURL: getEnv("DATABASE_URL", "postgres://root:root@localhost"),
		Port:        getEnv("PORT", "3030"),
	}
}

// unexported - lowercase
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}

	return defaultValue
}
