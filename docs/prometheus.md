# Prometheus Metrics

## Environment Variables

The following variables are supported for the Prometheus tracing/metrics gathering.

| ENV Var                     | Required | Description                                                                                              |
| --------------------------- | -------- | -------------------------------------------------------------------------------------------------------- |
| PROMETHEUS_SERVICE_PROTOCOL | No       | Indicate how we will communicate with Prometheus. Defaults to `http`                                     |
| PROMETHEUS_SERVICE_ADDR     | Yes*     | Indicate Prometheus URI. If set, enables Prometheus tracing when initPrometheusTracing(string) is called |

## Prometheus

 More info on Prometheus can be found [here](https://www.prometheus.io/)

### Quick Start (Local)
The following docker command will run the laster docker image for Prometheus

```sh
docker run -p 9090:9090 prom/prometheus
export PROMETHEUS_SERVICE_ADDR=localhost:9090
```

Then navigate to [http://localhost:9090](http://localhost:9090)

For more information on port usage, go [here](https://www.jaegertracing.io/docs/1.11/getting-started/)

Commonly used ports

| Port  | Protocol | Component | Function                                   |
| ----- | -------- | --------- | ------------------------------------------ |
| 9090  | HTTP     | query     | Serves Frontend                            |
| 14268 | HTTP     | collector | accept jaeger.thrift directly from clients |
