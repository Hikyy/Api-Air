package config

import (
	"fmt"
	"os"
	
	"github.com/joho/godotenv"
)

func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func Postgres() map[string]string {
	err := godotenv.Load()

	if err != nil {
		fmt.Println("Failed to load .env file:", err)
	}

	host := getEnv("POSTGRES_HOST", "localhost")
	port := getEnv("POSTGRES_PORT", "5432")
	dbuser := getEnv("POSTGRES_USER", "root")
	password := getEnv("POSTGRES_PASSWORD", "root")
	dbname := getEnv("POSTGRES_DB", "postgres")

	return map[string]string{
		"host":     host,
		"port":     port,
		"dbuser":   dbuser,
		"password": password,
		"dbname":   dbname,
	}
}
