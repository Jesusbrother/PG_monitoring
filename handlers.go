package main

import (
	"database/sql"
	"html/template"
	"net/http"
)

func dashboardHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		connections, err := getActiveConnections(db)
		if err != nil {
			http.Error(w, "Database query error (connections)", http.StatusInternalServerError)
			return
		}

		size, err := getDatabaseSize(db)
		if err != nil {
			http.Error(w, "Database query error (size)", http.StatusInternalServerError)
			return
		}

		sizeBytes, err := getDatabaseSizeBytes(db)
		if err != nil {
			http.Error(w, "Database query error (size bytes)", http.StatusInternalServerError)
			return
		}

		longRunning, err := getLongRunningQueriesCount(db)
		if err != nil {
			http.Error(w, "Database query error (long running queries)", http.StatusInternalServerError)
			return
		}

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
		}{
			Connections:        connections,
			DatabaseSize:       size,
			DatabaseSizeBytes:  sizeBytes,
			LongRunningQueries: longRunning,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Page rendering error", http.StatusInternalServerError)
			return
		}
	}
}
