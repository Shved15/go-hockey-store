package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

// Config represents the configuration parameters for the database connection.
type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	SSLMode    string
}

// DBConnect returns a string representing the database connection parameters.
// It uses the configuration values to construct a connection string.
func (c *Config) DBConnect() string {
	return fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%s",
		c.DBHost, c.DBPort, c.DBUser, c.DBName, c.DBPassword, c.SSLMode)
}

// GetConfig returns a new Config struct with default values set.
func GetConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	return Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     mustParseInt(os.Getenv("DB_PORT")),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		SSLMode:    os.Getenv("SSL_MODE"),
	}
}

// mustParseInt converts a string to an integer, throws an error if the conversion fails
func mustParseInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		log.Fatalf("Error parsing int: %v", err)
	}
	return i
}
