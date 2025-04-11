package metrics

import (
	"database/sql"
)

func GetActiveConnections(db *sql.DB) (int, error) {
	var count int
	err := db.QueryRow(`SELECT count(*) FROM pg_stat_activity;`).Scan(&count)
	return count, err
}
