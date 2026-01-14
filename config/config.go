package config

import "database/sql"

type AppConfig struct {
	Port           string
	DatabaseURL    string
	JWTSecret      string
	AdminEmail     string
	Environment    string
	MaxRequestSize int64
}

func Load() AppConfig {
	return AppConfig{}
}
