package main

import (
	"database/sql"
	"html/template"
	"net/http"

	"github.com/Jesusbrother/PG_monitoring/metrics"
)

func dashboardHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		connections, _ := metrics.GetActiveConnections(db)
		dbSize, _ := metrics.GetDatabaseSizePretty(db)
		dbSizeBytes, _ := metrics.GetDatabaseSizeBytes(db)
		longQueries, _ := metrics.GetLongRunningQueriesCount(db)
		containerSizeBytes, _ := metrics.GetContainerSize()

		tmpl, err := template.ParseFiles("templates/dashboard.html")
		if err != nil {
			http.Error(w, "Template rendering error", http.StatusInternalServerError)
			return
		}

		data := struct {
			Connections        int
			DatabaseSize       string
			DatabaseSizeBytes  int64
			LongRunningQueries int
			ContainerSizeBytes int64
		}{
			Connections:        connections,
			DatabaseSize:       dbSize,
			DatabaseSizeBytes:  dbSizeBytes,
			LongRunningQueries: longQueries,
			ContainerSizeBytes: containerSizeBytes,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Page rendering error", http.StatusInternalServerError)
			return
		}
	}
}
