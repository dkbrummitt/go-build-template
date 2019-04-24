package metrics

import (
	"errors"
	"os"

	"go.opencensus.io/exporter/prometheus"
	"go.opencensus.io/stats/view"
)

var (
	prometheusSvcAddr = os.Getenv("PROMETHEUS_SERVICE_ADDR")
	prometheusSvcPro  = os.Getenv("PROMETHEUS_SERVICE_PROTOCOL")
)

func initPrometheusTracing(ns string) error {
	var err error

	//check to see if feature is enabled
	if prometheusSvcAddr == "" {
		//Feature disabled, quit early
		return err
	}

	//validate we have a valid name space
	if ns == "" {
		return errors.New("A non-empty name-space is required to enable tracing")
	}

	if prometheusSvcPro == "" {
		prometheusSvcPro = defaultProtocol
	}

	//create the exporter
	exporter, err := prometheus.NewExporter(prometheus.Options{
		// Endpoint: fmt.Sprintf("%s://%s", prometheusSvcPro, prometheusSvcAddr),
		Namespace: ns,
	})
	if err != nil {
		return err
	}
	//register general view exporter
	view.RegisterExporter(exporter)

	return err
}

//Add view/stat?
func addViewStat() {}

//in code usage
//stats.Record(ctx, statName.M(int64(someIntValue))) // for simple count

//tracing
// ctx, span:= trace.StartSpan(stc,"/messages")
// defer span.End()
