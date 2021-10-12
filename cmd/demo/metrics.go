package main

import "github.com/prometheus/client_golang/prometheus"

var metricCounter prometheus.Counter
var metricGauge prometheus.Gauge
var metricHistogram prometheus.Histogram

func init() {
	namespace := "metricsdemo"
	subsystem := "demobin"
	metricCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "basic-counter",
		Help:      "Basic example of a counter",
	})
	metricGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "basic-gauge",
		Help:      "Basic example of a gauge",
	})
	metricHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "basic-histogram",
		Help:      "Basic example of a histogram",

		// This one adds a bucket field
		Buckets: prometheus.LinearBuckets(0, 50, 11), // 0, 50, 100, 150 ... 500
	})
}

func registerMetrics() {
	prometheus.MustRegister(
		metricCounter,
		metricGauge,
		metricHistogram,
	)
}
