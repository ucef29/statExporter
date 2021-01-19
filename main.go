package main

import (
	_ "errors"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type statCollector struct {
	fileSize *prometheus.Desc
}

func newStatCollector() *statCollector {
	return &statCollector{
		fileSize: prometheus.NewDesc("file_size", "Shows targeted file size", []string{"path"}, nil),
	}

}

func (collector *statCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.fileSize
}

func (collector *statCollector) Collect(ch chan<- prometheus.Metric) {
	file := "/tmp/file"
	size := getFileSize(file)

	ch <- prometheus.MustNewConstMetric(collector.fileSize, prometheus.GaugeValue, float64(size), file)
}

func getFileSize(file string) int64 {

	if _, err := os.Stat(file); os.IsNotExist(err) {
		return 0

	}

	if f, _ := os.Stat(file); f.Size() == 0 {
		return 0
	}

	info, _ := os.Stat(file)

	size := info.Size()

	return size
}

func main() {

	fileSize := newStatCollector()
	prometheus.MustRegister(fileSize)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
