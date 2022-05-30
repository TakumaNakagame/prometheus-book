package main

import (
	"flag"
	"log"
	"net/http"
	"fmt"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type myCollector struct{}

var (
	myCount = prometheus.NewCounter(prometheus.CounterOpts{
		Name:      "my_count",
		Help:      "my counter help",
	})
	countValue = float64(0)
)

func (c myCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- myCount.Desc()
}

func (c myCollector) Collect(ch chan<- prometheus.Metric) {
	
	ch <- prometheus.MustNewConstMetric(
		myCount.Desc(),
		prometheus.CounterValue,
		float64(countValue),
	)
}

func handler(w http.ResponseWriter, r *http.Request){
	countValue++
	fmt.Fprintf(w, "Hello :) %f", countValue)
}

var addr = flag.String("listen-address", ":8080", "The address to listen on for HTTP requests.")

func main() {
	flag.Parse()
	
	var c myCollector
	prometheus.MustRegister(c)
	
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(*addr, nil))
}