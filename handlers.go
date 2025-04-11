package main

import (
	"database/sql"
	"html/template"
	"net/http"
	"os/exec"
	"strconv"
	"strings"
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

		containerSizeBytes, err := getContainerSize()
		if err != nil {
			http.Error(w, "Container size error", http.StatusInternalServerError)
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
			ContainerSizeBytes int64
		}{
			Connections:        connections,
			DatabaseSize:       size,
			DatabaseSizeBytes:  sizeBytes,
			LongRunningQueries: longRunning,
			ContainerSizeBytes: containerSizeBytes,
		}

		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Page rendering error", http.StatusInternalServerError)
			return
		}
	}
}

func getContainerSize() (int64, error) {
	cmd := exec.Command("docker", "exec", "dockerpg", "du", "-sb", "/var/lib/postgresql/data")
	output, err := cmd.Output()
	if err != nil {
		return 0, err
	}

	parts := strings.Fields(string(output))
	if len(parts) < 1 {
		return 0, err
	}

	sizeBytes, err := strconv.ParseInt(parts[0], 10, 64)
	if err != nil {
		return 0, err
	}

	return sizeBytes, nil
}
