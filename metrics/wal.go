package metrics

import (
	"database/sql"
)

func GetWALSize(db *sql.DB) (int64, error) {
	var size int64
	err := db.QueryRow(`SELECT SUM(size) FROM pg_ls_waldir();`).Scan(&size)
	return size, err
}
