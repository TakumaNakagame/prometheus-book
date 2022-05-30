package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type myCollector struct{}

var (
	samleGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "sample_metrics",
		Help: "sample metrics from mysql",
	})
	metrics_value = float64(0)
)

func (c myCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- samleGauge.Desc()
}

func (c myCollector) Collect(ch chan<- prometheus.Metric) {

	ch <- prometheus.MustNewConstMetric(
		samleGauge.Desc(),
		prometheus.GaugeValue,
		float64(metrics_value),
	)
}

func get_metrics() {
	db, err := sql.Open("mysql", "root:prom@tcp(localhost:3306)/handson")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err)
	}

	err = db.QueryRow("select value from metrics where id=1 limit 1").Scan(&metrics_value)
	if err != nil {
		log.Fatal(err)
	}
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
	get_metrics()
	var c myCollector
	reg := prometheus.NewRegistry()
	reg.MustRegister(c)

	h := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)
	// var c myCollector
	// prometheus.MustRegister(c)
	
	// h := promhttp.Handler()
	// h.ServeHTTP(w, r)
}

var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")

func main() {
	flag.Parse()

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		metricsHandler(w, r)
	})
	log.Fatal(http.ListenAndServe(*addr, nil))
}
