---
marp: true
theme: gaia
_class: lead
paginate: true
backgroundColor: #fff
---

# Metrics

Prometheus 101

---

## Observability -- Asking new questions quickly

Three pillars of observability

1. Logging
    * log that we all know and love
1. Tracing
    * tracing the flow of events through a series of connected (micro-)services
1. **Metrics**
    * numeric representations of data _over time_.

---

## What is Monitoring and What are Metrics?

* Allows insight into how things are performing
* Alerts when something crashes
* Or even identify problems before it crashes

Examples of potential problems:

* memory at 70% and rising
* disk space filling up
* application latency increasing

---

Googles Site Reliability Engineering calls out 4 main metrics

* **latency**: time to service a request
* **traffic**: requests / second
* **error**: error rate of requests
* **saturation**: fullness of a service

---

## Prometheus as a monitoring platform

Prometheus is an open-source systems monitoring and alerting toolkit originally built at SoundCloud, started in 2012.

Prometheus collects and stores its metrics as time series data,
i.e. metrics information is stored with the timestamp at which it was recorded,
alongside optional key-value pairs called labels.

---

## Prometheus Features

Prometheus's main features are:

* a multi-dimensional data model with time series data identified by `metric` name and `key/value` pairs
* `PromQL`, a flexible query language to leverage this dimensionality
* no reliance on distributed storage; single server nodes are autonomous
* time series collection happens via a pull model over HTTP
* pushing time series is supported via an intermediary gateway
* targets are discovered via service discovery or static configuration
* multiple modes of graphing and dashboarding support

See [Prometheus overview docs](https://prometheus.io/docs/introduction/overview/) for more.

Note to self: We will focus on 1&2 today.  3/4/5 are nice to know about.  6 we might show if there is time.

Note: Prometheus should not be used 100% accuracy use cases, such as billing.

---

## What units are monitored?  These are called "Metrics"

A physical computer: `CPU usage`, `Memory usage`, `disk space used`, `disk read/write`, `fan speed`
web server: `request times`, `error rates`
database: `number of active connections`, `number of active queries`.

CubicSvr:

* Current recognizer count
* total bytes processed
* rtf over time
* error counts
