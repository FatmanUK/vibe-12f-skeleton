package metrics

import (
	"os"
	"log"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartMetricsServer() {
	port := os.Getenv("METRICS_SERVER_PORT")
	if port == "" {
		port = "9090"
	}
	log.Printf("Starting metrics server on :%s\n", port)
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start metrics server: %v", err)
	}
}
