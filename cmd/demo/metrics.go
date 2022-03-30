package main

import (
	"context"
	"errors"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func runMetricsServer(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())
	svr := http.Server{Addr: ":9850", Handler: mux}

	go func() {
		<-ctx.Done()
		svr.Shutdown(context.Background())
	}()

	err := svr.ListenAndServe()
	if err != nil && errors.Is(err, http.ErrServerClosed) {
		return nil // We expect this error on closing successfully
	}

	return err

}

var metricCounter prometheus.Counter
var metricGauge prometheus.Gauge
var metricHistogram prometheus.Histogram

func init() {
	namespace := "metricsdemo"
	subsystem := "demobin"
	metricCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "basiccounter",
		Help:      "Basic example of a counter",
	})
	metricGauge = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "basicgauge",
		Help:      "Basic example of a gauge",
	})
	metricHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Namespace: namespace,
		Subsystem: subsystem,
		Name:      "basichistogram",
		Help:      "Basic example of a histogram",

		// Histograms also have a filed for buckets a bucket field
		Buckets: prometheus.LinearBuckets(0.5, 0.1, 11), // 500, 600, 700, ... 1400, 1500 ms
	})
}

func registerMetrics() {
	prometheus.MustRegister(
		metricCounter,
		metricGauge,
		metricHistogram,
	)
}
