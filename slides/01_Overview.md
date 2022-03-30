---
marp: true
theme: gaia
_class: lead
# paginate: true
paginate: false
backgroundColor: #fff
---

# Metrics

Prometheus 101

---

## Observability 

#### Asking new questions quickly

Three pillars of observability

1. Logging
    * log that we all know and love
1. Tracing
    * tracing the flow of events through a series of connected (micro-)services
1. **Metrics**
    * numeric representations of data _over time_.

---

## What is monitoring?

* Allows insight into how things are performing
* Alerts when something crashes
* Or even identify problems before it crashes

Examples of potential problems:

* memory at 70% and rising
* disk space filling up
* application latency increasing

---

Googles Site Reliability Engineering calls out 4 main metrics to measure for any given service

* **latency**: time to service a request
* **traffic**: requests / second
* **error**: error rate of requests
* **saturation**: fullness of a service

---

## Prometheus as a monitoring platform

Prometheus is an open-source systems monitoring and alerting toolkit originally built at SoundCloud, started in 2012.


---

## Prometheus Features

Prometheus's main features are:

* a multi-dimensional data model with time series data identified by `metric` name and `key/value` pairs
* `PromQL`, a flexible query language to leverage this dimensionality
* no reliance on distributed storage; single server nodes are autonomous
* time series collection happens via a pull model over HTTP
* targets are discovered via service discovery or static configuration

See [Prometheus overview docs](https://prometheus.io/docs/introduction/overview/) for more.

Note: Prometheus should not be used 100% accuracy use cases, such as billing.

---

## What units are monitored?  

A physical computer: `CPU`, `RAM usage`, `disk space`, `fan speed`
web server: `request times`, `error rates`
database: `number of active connections`, `number of active queries`.

CubicSvr: `Current recognizer count`, `total bytes processed`, `rtf over time`, `error counts`
