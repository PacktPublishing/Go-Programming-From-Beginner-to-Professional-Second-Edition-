package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	healthzCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "healthz_calls_total",
		Help: "Total number of calls to the healthz endpoint.",
	})
)

func init() {
	prometheus.MustRegister(healthzCounter)
}

func main() {
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		healthzCounter.Inc()
		w.WriteHeader(http.StatusOK)
		fmt.Println("Monitoring endpoint invoked! Counter was incremented!")
	})
	http.Handle("/metrics", promhttp.Handler())
	server := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Server listening on port 8080...")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("Error starting server: %s\n", err)
	}
}
