package main

import (
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

var (
	httpRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "status"},
	)
)

func init() {
	prometheus.MustRegister(httpRequests)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	httpRequests.WithLabelValues(r.Method, "200").Inc()

	fmt.Fprintf(w, "Hello, World!")
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	promhttp.Handler().ServeHTTP(w, r)
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/metrics", metricsHandler)

	fmt.Println("Server is running at http://localhost:8080/")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
