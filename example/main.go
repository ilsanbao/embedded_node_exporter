package main

import (
	"log"
	"net/http"
	"ipip.net/embedded_node_exporter"
	"github.com/prometheus/client_golang/prometheus"
	"runtime"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {

	var addr = kingpin.Flag("addr", "").Default(":8080").String()

	kingpin.Parse() // embedded_node_exporter dependent

	handle := embedded_node_exporter.NewHandler(true, 20)

	test := prometheus.NewGauge(prometheus.GaugeOpts{
		Name:"test_goroutine_number",
		Help:"",
	})

	test.Set(float64(runtime.NumGoroutine()))

	handle.MustRegister(test)

	http.Handle("/metrics", handle)
	log.Fatal(http.ListenAndServe(*addr, nil))
}