package metrics

import (
	"errors"
	"fmt"
	"log"
	"os"

	"contrib.go.opencensus.io/exporter/jaeger"
	"go.opencensus.io/trace"
)

const (
	defaultJaegerAgentPort = "6831"
	defaultJaegerPort      = "14268"
)

var (
	jaegerSvcAddr   = os.Getenv("JAEGER_SERVICE_ADDR")
	jaegerSvcPro    = os.Getenv("JAEGER_SERVICE_PROTOCOL")
	jaegerSvcPort   = os.Getenv("JAEGER_SERVICE_PORT")
	jaegerAgentAddr = os.Getenv("JAEGER_AGENT_ADDR")
	jaegerAgentPro  = os.Getenv("JAEGER_AGENT_PROTOCOL")
	jaegerAgentPort = os.Getenv("JAEGER_AGENT_PORT")
)

func initJaegerTracing(opts MetricOptions) (err error) {

	// For tracing we have to do pretty much the same but using
	// the Jaeger exporter. We also have to tell the exporter
	// where it should send these traces:
	if jaegerSvcAddr == "" {
		//Feature disabled, quit early
		fmt.Println("WARNING Jaeger disabled")
		return
	}
	if opts.ServiceName == "" {
		fmt.Println("ERROR: Jaeger Missing/Empty ServiceName")
		return errors.New("A non-empty service name is required to enable tracing")
	}
	fmt.Println("Initializing Jaeger")

	// set defaults

	if jaegerSvcPro == "" {
		jaegerSvcPro = defaultProtocol
	}

	if jaegerSvcPort == "" {
		jaegerSvcPort = defaultJaegerPort
	}

	if jaegerAgentPort == "" {
		jaegerAgentPort = defaultJaegerAgentPort
	}

	// trace.ApplyConfig(trace.Config{
	// 	DefaultSampler: trace.AlwaysSample(),
	// })
	agentEndpointURI := "localhost:6831"
	collectorEndpointURI := "http://localhost:14268/api/traces"
	fmt.Println("\t Jaeger agentEndpointURI", agentEndpointURI)
	fmt.Println("\t Jaeger collectorEndpointURI", collectorEndpointURI)
	agentEndpointURI = fmt.Sprintf("%s:%s", jaegerAgentAddr, jaegerAgentPort)
	collectorEndpointURI = fmt.Sprintf("%s://%s:%s/api/traces", jaegerSvcPro, jaegerSvcAddr, jaegerSvcPort)
	fmt.Println("\t Jaeger agentEndpointURI*", agentEndpointURI)
	fmt.Println("\t Jaeger collectorEndpointURI*", collectorEndpointURI)

	jex, err := jaeger.NewExporter(jaeger.Options{
		AgentEndpoint:     agentEndpointURI,
		CollectorEndpoint: collectorEndpointURI,
		ServiceName:       opts.ServiceName,
	})
	if err != nil {
		log.Fatalf("Failed to create Jaeger exporter: %s", err.Error())
	}
	trace.RegisterExporter(jex)
	return
}
