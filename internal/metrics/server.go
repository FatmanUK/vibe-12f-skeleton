package metrics

import (
	"log"
	"net/http"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartMetricsServer() {
	log.Println("Starting metrics server on :9090")
	http.Handle("/metrics", promhttp.Handler())
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatalf("Failed to start metrics server: %v", err)
	}
}
