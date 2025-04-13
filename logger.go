package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"time"

	"github.com/Jesusbrother/PG_monitoring/metrics"
)

type MonitoringData struct {
	Connections        int   `json:"connections"`
	DatabaseSizeBytes  int64 `json:"database_size_bytes"`
	ContainerSizeBytes int64 `json:"container_size_bytes"`
	RAMUsed            int64 `json:"ram_used"`
	RAMTotal           int64 `json:"ram_total"`
	WALSize            int64 `json:"wal_size"`
	LongRunningQueries int   `json:"long_running_queries"`
}

// Background collector function
func collectAndInsertMetrics(db *sql.DB, intervalMinutes int) {
	connections, _ := metrics.GetActiveConnections(db)
	dbSizeBytes, _ := metrics.GetDatabaseSizeBytes(db)
	containerSizeBytes, _ := metrics.GetContainerSize()
	ramUsed, ramTotal, _ := metrics.GetContainerRAMUsage()
	walSize, _ := metrics.GetWALSize(db)
	longQueries, _ := metrics.GetLongRunningQueriesCount(db)

	data := MonitoringData{
		Connections:        connections,
		DatabaseSizeBytes:  dbSizeBytes,
		ContainerSizeBytes: containerSizeBytes,
		RAMUsed:            ramUsed,
		RAMTotal:           ramTotal,
		WALSize:            walSize,
		LongRunningQueries: longQueries,
	}

	checkAndInsertMonitoringLog(db, data, intervalMinutes)
}

func checkAndInsertMonitoringLog(db *sql.DB, data MonitoringData, intervalMinutes int) {
	var lastTime time.Time
	err := db.QueryRow(`SELECT collected_at FROM monitoring_logs ORDER BY collected_at DESC LIMIT 1`).Scan(&lastTime)

	if err != nil && err != sql.ErrNoRows {
		log.Printf("Error checking last monitoring log: %v", err)
		return
	}

	if time.Since(lastTime) < time.Duration(intervalMinutes)*time.Minute {
		log.Println("Less than configured interval since last log, skipping insert.")
		return
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("Error marshaling monitoring data to JSON: %v", err)
		return
	}

	_, err = db.Exec(`INSERT INTO monitoring_logs (data) VALUES ($1)`, jsonData)
	if err != nil {
		log.Printf("Error inserting monitoring log: %v", err)
	} else {
		log.Println("Monitoring log inserted successfully.")
	}
}
