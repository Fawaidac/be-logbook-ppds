package configs

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	JWTSecret  string
}

func LoadConfig() *Config {
	_ = godotenv.Load()

	return &Config{
		ServerPort: getEnv("SERVER_PORT", "8080"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "postgres"),
		DBPassword: getEnv("DB_PASSWORD", "fawaidach11"),
		DBName:     getEnv("DB_NAME", "logbook-ppds"),
		DBSSLMode:  getEnv("DB_SSLMODE", "disable"),
		JWTSecret:  getEnv("JWT_SECRET", "super-secret-key-12345"),
	}
}

func getEnv(key, defaultValue string) string {
	if val, exists := os.LookupEnv(key); exists && val != "" {
		return val
	}
	return defaultValue
}

