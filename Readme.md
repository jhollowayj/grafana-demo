# Grafana demonstration

Example Go program to demonstrate using prometheus, loki, and grafana for observability.

Items of importance:

    * Running multiple instances of a service
    * graphing metrics (line graph, gauge)
    * vectors of those metrics
    * logs
    * graphs of the logs
    * templates/repeats (what's the technical name?)

## File directory

`cmd` - hosts the go code
`bin` - hosts the generated go binary
`deploy` - hosts the docker-compose and all relevant config files.
`docs` - hosts graphviz source files and generated images.

## Overview of the go program

TODO

## Overview of the docker-compose setup

TODO

## Overview of the stacks

Prometheus scrapes your application ever {15s}, then Grafana queries Prometheus as needed.
Promtail watches file and pushes logs to Loki, then Grafana queries Loki as needed.
Promtail is typically a side-car to your app (same pod, different container)

## Important pages

* http://localhost:9090/targets

## Prometheus Metric Types

Docs: https://prometheus.io/docs/concepts/metric_types/

Counters - Only goes up. (Use case: number of requests served, tasks completed, or errors)
Gauges - Can go up or down.  (Use case: temperatures, memory usage, concurrent requests)
Histograms - Bucketized counts (Use case: request durations, response sizes; help determine SLO (i.e. serve 95% of requests within 300ms))
Summaries - Like histograms, but more complicated.  See https://prometheus.io/docs/practices/histograms/ for more.

## PromQL

Docs: https://prometheus.io/docs/prometheus/latest/querying/basics/

## LogQL

Docs: https://grafana.com/docs/loki/latest/logql/
