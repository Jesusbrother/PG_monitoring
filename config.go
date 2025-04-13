package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost             string
	DBPort             string
	DBUser             string
	DBPassword         string
	DBName             string
	SSLMode            string
	LogIntervalMinutes int
}

func loadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Failed to load .env file: %v", err)
	}

	logIntervalStr := os.Getenv("LOG_INTERVAL_MINUTES")
	logInterval, err := strconv.Atoi(logIntervalStr)
	if err != nil {
		log.Printf("Invalid LOG_INTERVAL_MINUTES, using default 30 minutes")
		logInterval = 30
	}

	return Config{
		DBHost:             os.Getenv("DB_HOST"),
		DBPort:             os.Getenv("DB_PORT"),
		DBUser:             os.Getenv("DB_USER"),
		DBPassword:         os.Getenv("DB_PASSWORD"),
		DBName:             os.Getenv("DB_NAME"),
		SSLMode:            os.Getenv("SSL_MODE"),
		LogIntervalMinutes: logInterval,
	}
}
