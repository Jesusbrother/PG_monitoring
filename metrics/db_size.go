package metrics

import (
	"database/sql"
)

func GetDatabaseSizePretty(db *sql.DB) (string, error) {
	var size string
	err := db.QueryRow(`SELECT pg_size_pretty(pg_database_size(current_database()));`).Scan(&size)
	return size, err
}

func GetDatabaseSizeBytes(db *sql.DB) (int64, error) {
	var sizeBytes int64
	err := db.QueryRow(`SELECT pg_database_size(current_database());`).Scan(&sizeBytes)
	return sizeBytes, err
}
