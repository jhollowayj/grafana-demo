---
marp: true
theme: gaia
# math: katex
---

# Prometheus Concepts

---

## Data Model

---

### Data Model: Metric `names` and `labels`

Each time series is a unique combination of _metric name_ and optional key-value pairs called _labels_.

Example name: `http_request_total`
Example labels: `{method='get', status=200, handler='/version'}`

<br> 

Note: Each combination of label values is a new time series.  Be careful to keep cardinality down.
i.e. Using `http status` is ok.  Using `User IDs` is not.

---

### Data Model: Samples

Each sample consists of

* a float64 value
* a millisecond-precision timestamp

---

### Data Model: Notation

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

### Metric Types: Counter

Cumulative metric, represents monotonically increasing counter.
Can only increase or be reset to zero on restart.

Used for:

* requests served
* tasks completed
* error counts

---

### Metric Types: Gauges

Single numeric value that can go up and down

Used for:

* temperatures
* current memory usage
* number of concurrent requests

---

### Metric Types: Histogram

A histogram samples observations, counts them in configurable buckets.
Also provides sum of all observed values.

Each base metric name exposes multiple series:

* counters for each bucket, exposed as `<basename>_bucket{le="<upper inclusive bound>"}`
* total sum of all observed values, exposed as `<basename>_sum`
* count of events, exposed as `<basename>_count` (identical to `<basename>_bucket{le="+Inf"}`)

Use of `histogram_quantile()` _function_ to calculate quantiles.
i.e. What was the 99th quantile for response latency?

<!--
Can be used for Apdex ("Application Performance Index") calculations:
$$
    Apdex_t = \dfrac{SatisfiedCount + (0.5 * ToleratingCount) + (0 * FrustratedCount)}{TotalSamples}
$$
-->
---

### Metric Types: Summary

Very similar to histogram.
Calculate quantiles over a sliding time window.
Doesn't need pre-defined buckets, but comes with a much higher computation cost instead.

---

## Jobs and Instances

`job`: The configured job name that the target belongs to.
`instance`: The `<host>:<port>` part of the target's URL that was scraped.

Example of web server with 4 replication instances:

```txt
    job: api-server
        instance 1: 1.2.3.4:5670
        instance 2: 1.2.3.4:5671
        instance 3: 5.6.7.8:5670
        instance 4: 5.6.7.8:5671
```
