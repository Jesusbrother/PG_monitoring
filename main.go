package main

import (
	"log"
	"net/http"
)

func main() {
	cfg := loadConfig()
	db := initDB(cfg)
	defer db.Close()

	http.HandleFunc("/", dashboardHandler(db))

	log.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
