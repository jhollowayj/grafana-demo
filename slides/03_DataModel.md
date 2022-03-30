---
marp: true
theme: default
---

# Data Model

## Metric Names and labels

Each time series is a unique _metric name_ and optional _labels_ as key-value pairs.

`http_request_total` would be the name.

One set of labels would be `{method=put}`, or `{method=put, status=200}`

Note: each combination of metric name + labels is a new time series.  Don't go crazy with labels.

## Samples

Each sample consists of

* a float64 value
* a millisecond-precision timestamp

---

## Notation

```txt
<metric name>{<label name>=<label value>, ...}
```

```txt
api_http_requests_total{method="POST", handler="/messages"}
```
