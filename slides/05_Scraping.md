---
marp: true
theme: gaia
---

# Scraping

## Serve `/metrics` endpoint

```go
import "github.com/prometheus/client_golang/prometheus/promhttp"

http.Handle("/metrics", promhttp.Handler())
http.ListenAndServe("localhost:1234", nil)
```

---

## Configure Prometheus to scrape our target

`prometheus.yaml`:

```yaml
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'MyGoService'
    static_configs:
      - targets: ['localhost:1234']
```
