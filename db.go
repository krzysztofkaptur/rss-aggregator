package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/krzysztofkaptur/rss-aggregator/internal/database"
)

func InitDB() (*database.Queries, error) {
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbSslMode := os.Getenv("DB_SSL_MODE")

	connStr := fmt.Sprintf("user=%v dbname=%v password=%v sslmode=%v", dbUser, dbName, dbPassword, dbSslMode)
	conn, dbErr := sql.Open("postgres", connStr)

	if dbErr != nil {
		return nil, dbErr
	}

	store := database.New(conn)

	return store, nil
}