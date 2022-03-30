---
marp: true
theme: default
---

# Instrumenting Go services

## Import client library

```go
import "github.com/prometheus/client_golang/prometheus"
```

## Adding http endpoint

```go
import "net/http"
import "github.com/prometheus/client_golang/prometheus/promhttp"

http.Handle("/metrics", promhttp.Handler())
http.ListenAndServe(*addr, nil)
```

## Create Metrics

```go
// Create single metrics
m = prometheus.NewCounter(prometheus.CounterOpts{})
m = prometheus.NewGauge(prometheus.GaugeOpts{})
m = prometheus.NewHistogram(prometheus.HistogramOpts{})
m = prometheus.NewSummary(prometheus.SummaryOpts{})

// Create a vector of metrics, with labels
m = prometheus.NewCounterVec(prometheus.CounterOpts{}, []string{"labels"})
m = prometheus.NewGaugeVec(prometheus.GaugeOpts{}, []string{"labels"})
m = prometheus.NewHistogramVec(prometheus.HistogramOpts{}, []string{"labels"})
m = prometheus.NewSummaryVec(prometheus.SummaryOpts{}, []string{"labels"})

// There's also *Func variants that let you us a callback.
// The function is called when scraped, so you can look up the value when needed.
// This is useful if you need to keep track of the count outside of metrics.
m = prometheus.NewCounterFunc(prometheus.CounterOpts{}, func() float64 {
    return curVal // Return the current value
})
```

### Options

```go
type Opts struct {
	Namespace string // optional
	Subsystem string // optional
	Name string // Required
	Help string // Explanation of what this metric tracks.
}

type GaugeOpts Opts
type CounterOpts Opts
type HistogramOpts struct {
	Opts
	Buckets []float64
}
```

Fully Qualified metric names are generated as `{Namespace}_{Subsystem}_{Name}`.
Note: names must be globally unique.

## Register the metrics

```go
prometheus.MustRegister(metricCounter, metricGauge, metricHistogram) // Will panic
// OR
err := prometheus.Register(metric)
```

Alternatively, you can use promauto to both create and register.  Beware panics though.

```go
import "github.com/prometheus/client_golang/prometheus/promauto"

promauto.NewCounter(prometheus.CounterOpts{})
```

## Making observations

```go
counter.Inc()   // +1
counter.Add(n) // +n; Must be >0

gauge.Set(n) // Jump right to n
gauge.Inc()  // +1
gauge.Dec()  // -1
gauge.Add(n) // +n (Can be negative)
gauge.Sub(n) // -n

histogram.Observe(n)

summary.Observe(n)
```
