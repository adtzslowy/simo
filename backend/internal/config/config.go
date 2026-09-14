package config

import "os"

type Config struct {
	AppEnv string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string

	ServerPort string
}

func Load() Config {
	return Config{
		AppEnv: getEnv("APP_ENV", "development"),

		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("DB_USER", "orbit"),
		DBPassword: getEnv("DB_PASSWORD", "orbit"),
		DBName:     getEnv("DB_NAME", "orbit"),

		ServerPort: getEnv("SERVER_PORT", "8080"),
	}
}

func getEnv(key, fallback string) string {
	value := os.Getenv(key)

	if value == "" {
		return fallback
	}

	return value
}
