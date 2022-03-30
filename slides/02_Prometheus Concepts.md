---
marp: true
theme: default
math: katex
---

# Prometheus Concepts

---

## Data Model

### Metric Names and labels

Each time series is a unique _metric name_ and optional _labels_ as key-value pairs.

`http_request_total` would be the name.

One set of labels would be `{method=put}`, or `{method=put, status=200}`

Note: each combination of metric name + labels is a new time series.  Don't go crazy with labels.

### Samples

Each sample consists of

* a float64 value
* a millisecond-precision timestamp

---

### Notation

```txt
<metric name>{<label name>=<label value>, ...}
```

```txt
api_http_requests_total{method="POST", handler="/messages"}
```

---

## Metric Types

4 Prometheus metric types:

* Counter
* Gauges
* Histogram
* Summary

---

### Counter

Cumulative metric, represents monotonically increasing counter.
Can only increase or be reset to zero on restart.

Used for:

* Requests served
* tasks completed
* error counts

---

### Gauges

Single numeric value that can go up and down

Used for:

* Temperatures
* current memory usage
* number of concurrent requests

---

### Histogram

A histogram samples observations, counts them in configurable buckets.
Also provides sum of all observed values.

Each base metric name exposes multiple series:

* counters for each bucket, exposed as `<basename>_bucket{le="<upper inclusive bound>"}`
* total sum of all observed values, exposed as `<basename>_sum`
* count of events, exposed as `<basename>_count` (identical to `<basename>_bucket{le="+Inf"}`)

Use of `histogram_quantile()` _function_ to calculate quantiles.
Can be used for Apdex ("Application Performance Index") calculations:

$$
    Apdex_t = \dfrac{SatisfiedCount + (0.5 * ToleratingCount) + (0 * FrustratedCount)}{TotalSamples}
$$

---

### Summary

Very similar to histogram. Calculates quantiles over a sliding time window.

---

## Jobs and Instances

`job`: The configured job name that the target belongs to.
`instance`: The `<host>:<port>` part of the target's URL that was scraped.

example of web server with 4x replications:

```txt
    job: api-server
        instance 1: 1.2.3.4:5670
        instance 2: 1.2.3.4:5671
        instance 3: 5.6.7.8:5670
        instance 4: 5.6.7.8:5671
```

---
