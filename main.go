package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	cfg := loadConfig()
	db := initDB(cfg)
	defer db.Close()

	// Start background monitoring goroutine
	go func() {
		for {
			collectAndInsertMetrics(db, cfg.LogIntervalMinutes)
			time.Sleep(time.Duration(cfg.LogIntervalMinutes) * time.Minute) // hardcode now (#todo time.Ticker)
		}
	}()

	http.HandleFunc("/", dashboardHandler(db))

	log.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
