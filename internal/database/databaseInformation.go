package database

import (
	"App/internal/config"
	"fmt"
)

// BuildConnectionString construit la chaîne de connexion à la base de données PostgreSQL
func BuildConnectionString() string {

	configPostgres := config.Postgres()

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		configPostgres["host"], configPostgres["port"], configPostgres["dbuser"], configPostgres["password"], configPostgres["dbname"])
}
