package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func initDB(cfg Config) *sql.DB {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.SSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Database is not responding: %v", err)
	}

	log.Println("Database connected successfully")
	return db
}

func getActiveConnections(db *sql.DB) (int, error) {
	var count int
	err := db.QueryRow(`SELECT count(*) FROM pg_stat_activity;`).Scan(&count)
	return count, err
}

func getDatabaseSize(db *sql.DB) (string, error) {
	var size string
	err := db.QueryRow(`SELECT pg_size_pretty(pg_database_size(current_database()));`).Scan(&size)
	return size, err
}

func getDatabaseSizeBytes(db *sql.DB) (int64, error) {
	var sizeBytes int64
	err := db.QueryRow(`SELECT pg_database_size(current_database());`).Scan(&sizeBytes)
	return sizeBytes, err
}

func getLongRunningQueriesCount(db *sql.DB) (int, error) {
	var count int
	err := db.QueryRow(`SELECT count(*) FROM pg_stat_activity WHERE state = 'active' AND now() - query_start > interval '5 seconds';`).Scan(&count)
	return count, err
}
