package main

import (
	"log"
	"net/http"
)

func main() {
	// load config
	cfg := loadConfig()

	// db connection
	db := initDB(cfg)
	defer db.Close()

	// routing
	http.HandleFunc("/", dashboardHandler(db))

	log.Println("Server is running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
