package metrics

import (
	"database/sql"
)

func GetLongRunningQueriesCount(db *sql.DB) (int, error) {
	var count int
	err := db.QueryRow(`SELECT count(*) FROM pg_stat_activity WHERE state = 'active' AND now() - query_start > interval '5 seconds';`).Scan(&count)
	return count, err
}
