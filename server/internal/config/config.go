package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort int
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     int
	DBName     string

	DBConnectionString string
}

func Load() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, proceeding with environment variables")
	}

	serverPort := getEnvAsInt("SERVER_PORT", 3000)
	dbUser := getEnv("DB_USER", "user")
	dbPassword := getEnv("DB_PASSWORD", "password")
	dbHost := getEnv("DB_HOST", "localhost")
	dbPort := getEnvAsInt("DB_PORT", 5432)
	dbName := getEnv("DB_NAME", "database")

	dbConnectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, strconv.Itoa(dbPort), dbName)

	cfg := &Config{
		ServerPort:         serverPort,
		DBUser:             dbUser,
		DBPassword:         dbPassword,
		DBHost:             dbHost,
		DBPort:             dbPort,
		DBName:             dbName,
		DBConnectionString: dbConnectionString,
	}

	if err := validateConfig(cfg); err != nil {
		return nil, fmt.Errorf("config validation failed: %w", err)
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	log.Printf("Environment variable %s not set, using fallback: %s", key, fallback)

	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	valueStr := getEnv(key, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	log.Printf("Environment variable %s not an integer or not set, using fallback: %d", key, fallback)

	return fallback
}

func validateConfig(cfg *Config) error {
	return nil
}
