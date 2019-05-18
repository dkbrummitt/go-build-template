package metrics

import (
	"errors"
	"fmt"
	"os"

	"contrib.go.opencensus.io/exporter/prometheus"
	"go.opencensus.io/stats/view"
)

var (
	prometheusSvcAddr = os.Getenv("PROMETHEUS_SERVICE_ADDR")
	prometheusSvcPro  = os.Getenv("PROMETHEUS_SERVICE_PROTOCOL")
)

func initPrometheusTracing(opts MetricOptions) error {
	// To export metrics for Prometheus, we have to create a
	// Prometheus exporter, attach it to stats view, and register
	// the exporter with the HTTP request muxer:

	var err error

	//check to see if feature is enabled
	if prometheusSvcAddr == "" {
		//Feature disabled, quit early
		return err
	}

	//validate we have a valid name space
	if opts.ServiceName == "" {
		return errors.New("A non-empty name-space is required to enable tracing")
	}
	fmt.Println("Initializing Prometheus")

	if prometheusSvcPro == "" {
		prometheusSvcPro = defaultProtocol
	}

	//create the exporter
	exporter, err := prometheus.NewExporter(prometheus.Options{
		// Endpoint: fmt.Sprintf("%s://%s", prometheusSvcPro, prometheusSvcAddr),
		Namespace: opts.ServiceName,
	})
	if err != nil {
		return err
	}
	//register general view exporter
	view.RegisterExporter(exporter)
	opts.Mux.Handle("/metrics", exporter)

	return err
}
