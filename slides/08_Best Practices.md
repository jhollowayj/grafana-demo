---
marp: true
theme: gaia
---

# Best Practices

---

## Metric Names

A metric name...

* should have a (single-word) application prefix relevant to the domain the metric belongs to. The prefix is sometimes referred to as namespace by client libraries. For metrics specific to an application, the prefix is usually the application name itself. Sometimes, however, metrics are more generic, like standardized metrics exported by client libraries. Examples:
  * **`prometheus`**`_notifications_total` (specific to the Prometheus server)
  * **`process`**`_cpu_seconds_total` (exported by many client libraries)
  * **`http`**`_request_duration_seconds` (for all HTTP requests)
* must have a single unit (i.e. do not mix seconds with milliseconds, or seconds with bytes).
* should have a suffix describing the unit, in plural form.
  * Note that an accumulating count has total as a suffix, in addition to the unit if applicable.
    * `http_request_duration`**`_seconds`**
    * `node_memory_usage`**`_bytes`**
    * `http_requests`**`_total`** (for a unit-less accumulating count)
    * `process_cpu_seconds`**`_total`** (for an accumulating count with unit)
    * `foobar_build`**`_info`** (for a pseudo-metric that provides metadata about the running binary)
* should represent the same logical thing-being-measured across all label dimensions.
  * request duration
  * bytes of data transfer
  * instantaneous resource usage as a percentage

## Labels

Use labels to differentiate the characteristics of the thing that is being measured:

`api_http_requests_total` - differentiate request types: `operation="create|update|delete"`
`api_request_duration_seconds` - differentiate request stages: `stage="extract|transform|load"`

Do not put the label names in the metric name, as this introduces redundancy and will cause confusion if the respective labels are aggregated away.

## Base Units

Always use base units.

| Family      | Base unit |
| Time        | seconds   |
| Temperature | celsius   |
| Length      | meters    |
| Bytes       | bytes     |
| Bits        | bytes     |
| Percent     | ratio     |

### Avoid missing metrics

Time series that are not present until something happens.
You can record a default value of `0` to force it's creation
Metrics with multiple labels need this several times if you want to populate them all.
