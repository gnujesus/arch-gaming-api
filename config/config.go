package config

type AppConfig struct {
	Port           string
	DatabaseURL    string
	JWTSecret      string
	AdminEmail     string
	Environment    string
	MaxRequestSize int64
}
