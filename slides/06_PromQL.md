---
marp: true
theme: gaia
---

# PromQL

---

## Expression language data types

In Prometheus's expression language, an expression or sub-expression can evaluate to one of four types:

**Scalar** - a simple numeric floating point value
**Instant vector** - a set of time series containing a single sample for each time series, all sharing the same timestamp
**Range vector** - a set of time series containing a range of data points over time for each time series

Query for an instant vector:
`http_requests_total{job="prometheus",group="canary"}`

Query for a range vector:
`http_requests_total{job="prometheus"}[5m]`

---

## String matching

* `=`: Select labels that are exactly equal to the provided string.
* `!=`: Select labels that are not equal to the provided string.
* `=~`: Select labels that regex-match the provided string. (Uses RE2 syntax)
* `!~`: Select labels that do not regex-match the provided string.

Example: `http_requests_total{environment=~"staging|testing|development",method!="GET"}`

---

## Time Durations

* `ms` - milliseconds
* `s` - seconds
* `m` - minutes
* `h` - hours
* `d` - days - assuming a day has always 24h
* `w` - weeks - assuming a week has always 7d
* `y` - years - assuming a year has always 365d

Examples:

```txt
5h
1h30m
5m
10s
```

---

## Operators

---

### Arithmetic binary operators

* `+` (addition)
* `-` (subtraction)
* `*` (multiplication)
* `/` (division)
* `%` (modulo)
* `^` (power/exponentiation)

Valid between:

* two scalars
* instant vector and a scalar
* two instant vectors

---

### Comparison binary operators

* `==` (equal)
* `!=` (not-equal)
* `>` (greater-than)
* `<` (less-than)
* `>=` (greater-or-equal)
* `<=` (less-or-equal)

Valid between:

* two scalars
* instant vector and a scalar
* two instant vectors

---

### Logical/set binary operators

* `and` (intersection)
* `or` (union)
* `unless` (complement)

---

<style scoped>
ul {
  font-size: 90%;
}
</style>
### Aggregation operators

Used to aggregate the elements of a single instant vector, resulting in a new vector of fewer elements with aggregated values.

* `sum` (calculate sum over dimensions)
* `min` (select minimum over dimensions)
* `max` (select maximum over dimensions)
* `avg` (calculate the average over dimensions)
* `group` (all values in the resulting vector are 1)
* `stddev` (calculate population standard deviation over dimensions)
* `stdvar` (calculate population standard variance over dimensions)
* `count` (count number of elements in the vector)
* `count_values` (count number of elements with the same value)
* `bottomk` (smallest k elements by sample value)
* `topk` (largest k elements by sample value)
* `quantile` (calculate φ-quantile (0 ≤ φ ≤ 1) over dimensions)

---

### Aggregation operator examples

Given the time series of `http_requests_total` with labels for `application`, `instance`, and `group`

```promql
# These two are the same:
# Keep labels `applications` and `groups`, collapse down the instance labels
sum without (instance) (http_requests_total)
sum by (application, group) (http_requests_total)

# total of HTTP requests seen in _all_ applications
sum(http_requests_total)

# count the number of binaries running each build version
count_values("version", build_version)

# get the 5 largest HTTP requests counts across all instances
topk(5, http_requests_total)
```

---

## Functions

<style scoped>
table {
  font-size: 80%;
}
</style>
Here's a full list.

|                    |                 |                      |                        |
| ------------------ | --------------- | -------------------- | ---------------------- |
| `abs()`            | `absent()`      | `absent_over_time()` | `ceil()`               |
| `changes()`        | `clamp()`       | `clamp_max()`        | `clamp_min()`          |
| `day_of_month()`   | `day_of_week()` | `days_in_month()`    | `delta()`              |
| `deriv()`          | `exp()`         | `floor()`            | `histogram_quantile()` |
| `holt_winters()`   | `hour()`        | `idelta()`           | `increase()`           |
| `irate()`          | `label_join()`  | `label_replace()`    | `ln()`                 |
| `log2()`           | `log10()`       | `minute()`           | `month()`              |
| `predict_linear()` | `rate()`        | `resets()`           | `round()`              |
| `scalar()`         | `sgn()`         | `sort()`             | `sort_desc()`          |
| `sqrt()`           | `time()`        | `timestamp()`        | `vector()`             |
| `year()`           |                 |                      |                        |

`<aggregation>_over_time()`: `sum`, `min`, `max`, `avg`, `stddev`, `stdvar`, `count`, `quantile`

`Trigonometric Functions`: `acos`, `acosh`, `asin`, `asinh`, `atan`, `atanh`, `cos`, `cosh`, `sin`, `sinh`, `tan`, `tanh`

---

## Query Examples

---
### Simple query examples

```promql
# all time series with the metric `http_requests_total`
http_requests_total 

# all time series with the metric `http_requests_total` and the given `job` and `handler` labels
http_requests_total{job="apiserver", handler="/api/comments"}

# a whole range of time (in this case 5 minutes) (range vector)
http_requests_total{job="apiserver", handler="/api/comments"}[5m] 

# all jobs that end with `server`
http_requests_total{job=~".*server"}

# All HTTP status codes except 4xx ones
http_requests_total{status!~"4.."}
```

---

### Examples using query functions

```promql
# per-second rate, as measured over the last 5 minutes
rate(http_requests_total[5m])

# If http_requests_total has `job` and `instance` labels, we can drop one of those with `sum by ()`
sum by (job) (
  rate(http_requests_total[5m])
)

# We can take the diff between two, as long as labels all match up
(instance_memory_limit_bytes - instance_memory_usage_bytes) / 1024 / 1024

# Then we could sum by application
sum by (app, proc) (
  instance_memory_limit_bytes - instance_memory_usage_bytes
) / 1024 / 1024
```

---

### Examples using query functions (cont)

Lets say we have these series:

```txt
instance_cpu_time_ns{app="lion",     proc="web",    rev="34d0f99", env="prod", job="cluster-manager"}
instance_cpu_time_ns{app="elephant", proc="worker", rev="34d0f99", env="prod", job="cluster-manager"}
instance_cpu_time_ns{app="turtle",   proc="api",    rev="4d3a513", env="prod", job="cluster-manager"}
instance_cpu_time_ns{app="fox",      proc="widget", rev="4d3a513", env="prod", job="cluster-manager"}
...
```

```promql
# get the top 3 CPU users grouped by application (app) and process type (proc)
topk(3, sum by (app, proc) (rate(instance_cpu_time_ns[5m])))

# count the number of running instances per application like this
count by (app) (instance_cpu_time_ns)
```
