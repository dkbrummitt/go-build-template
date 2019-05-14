package metrics

import (
	"errors"
	"fmt"
	"os"

	"contrib.go.opencensus.io/exporter/jaeger"
	"go.opencensus.io/stats/view"
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

type JaegerTracer struct {
	Exporter *jaeger.Exporter
}

func initJaegerTracing(sn string) (JaegerTracer, error) {
	var err error
	jt := JaegerTracer{}
	//check to see if feature is enabled
	if jaegerSvcAddr == "" {
		//Feature disabled, quit early
		return jt, err
	}

	//validate we have a valid service name
	if sn == "" {
		return jt, errors.New("A non-empty service name is required to enable tracing")
	}

	if jaegerSvcPro == "" {
		jaegerSvcPro = defaultProtocol
	}

	if jaegerSvcPort == "" {
		jaegerSvcPort = defaultJaegerPort
	}

	if jaegerAgentPort == "" {
		jaegerAgentPort = defaultJaegerAgentPort
	}

	agentEndpointURI := fmt.Sprintf("%s:%s", jaegerAgentAddr, jaegerAgentPort)
	collectorEndpointURI := fmt.Sprintf("%s://%s:%s/api/traces", jaegerSvcPro, jaegerSvcAddr, jaegerSvcPort)

	//create the exporter
	exporter, err := jaeger.NewExporter(
		jaeger.Options{
			AgentEndpoint:     agentEndpointURI,
			CollectorEndpoint: collectorEndpointURI,
			// Endpoint:          fmt.Sprintf("%s://%s:%s", jaegerSvcPro, jaegerSvcAddr, jaegerSvcPort),
			Process: jaeger.Process{
				ServiceName: sn,
			},
		})
	if err != nil {
		return jt, err
	}

	//register it
	trace.RegisterExporter(exporter)
	jt = JaegerTracer{
		Exporter: exporter,
	}
	return jt, err
}

func (jt *JaegerTracer) RegisterViews(views ...view.View) {
	for _, v := range views {
		//TODO how to load views to existing exporter on init...
		//or after the fact?
		//view.RegisterExporter(&jt.Exporter)
		fmt.Println("View:", v)
	}
}
