package main

import (
	"log"
	"twelve-factor-app/internal/metrics"
)

func initDatabase() error {
	// Stub function to connect to a database
	log.Println("Initializing database connection (stub)")
	return nil
}

func main() {
	err := initDatabase()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	metrics.StartMetricsServer()
}
