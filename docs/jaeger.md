# Jaeger Metrics

## Environment Variables

The following variables are supported for the Jaeger tracing.

| ENV Var                 | Required | Description                                                                                  |
| ----------------------- | -------- | -------------------------------------------------------------------------------------------- |
| JAEGER_SERVICE_PROTOCOL | No       | Indicate how we will communicate with Jaeger. Defaults to `http`                             |
| JAEGER_SERVICE_ADDR     | Yes*     | Indicate Jaeger URI. If set, enables Jaeger tracing when initJaegerTracing(string) is called |

## Jaeger

 More info on Jaeger can be found [here](https://www.jaegertracing.io/)

### Quick Start (Local)
The following docker command will run the laster docker image for Jaeger

```sh
docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14268:14268 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.11
export JAEGER_SERVICE_ADDR=localhost:16686
```
Then navigate to [http://localhost:16686](http://localhost:16686)

For more information on port usage, go [here](https://www.jaegertracing.io/docs/1.11/getting-started/)

Commonly used ports

| Port  | Protocol | Component | Function                                   |
| ----- | -------- | --------- | ------------------------------------------ |
| 16686 | HTTP     | query     | Serves Frontend                            |
| 14268 | HTTP     | collector | accept jaeger.thrift directly from clients |
