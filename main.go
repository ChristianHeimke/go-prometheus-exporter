package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"time"
)

func recordMetrics() {
	// go routine zum simulierten hochzählen des Event Counters alle 2 Sekunden
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()

}

var (
	// Erstellen und Hinzufügen einer neuen Counter Metric
	// (Simulierter Zähler für verarbeitete Events)
	opsProcessed = promauto.NewCounter((prometheus.CounterOpts{
		Name: "goapp_processed_ops_total",
		Help: "The total number of processed events",
	}))
)

func main() {
	// individuelle Metric aufzeichnen
	recordMetrics()

	// Metric unter dem Pfad bereitstellen zum abrufen
	http.Handle("/metrics", promhttp.Handler())

	// http-server unter dem Port 9123 verfügbar machen
	http.ListenAndServe(":9123", nil)
}
