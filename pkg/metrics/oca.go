package metrics

import (
	"errors"
	"fmt"
	"os"

	"go.opencensus.io/stats/view"

	"contrib.go.opencensus.io/exporter/ocagent"
	"go.opencensus.io/trace"
)

const (
	defaultServicePort = "55678"
)

var (
	ocaSvcAddr = os.Getenv("OPEN_CENSUS_AGENT_SERVICE_ADDR")
	ocaSvcPro  = os.Getenv("OPEN_CENSUS_AGENT_SERVICE_PROTOCOL")
	ocaSvcPort = os.Getenv("OPEN_CENSUS_AGENT_SERVICE_PORT")
)

func initOCAgent(sn string) error {
	var err error

	if ocaSvcAddr == "" {
		//quit early. feature disabled
		return err
	}
	if sn == "" {
		return errors.New("A non-empty service name is required to enable open census agent support")
	}

	if ocaSvcPro == "" {
		ocaSvcPro = defaultProtocol
	}
	if ocaSvcPort == "" {
		ocaSvcPort = defaultServicePort
	}

	oce, err := ocagent.NewExporter(
		ocagent.WithInsecure(),
		ocagent.WithServiceName(fmt.Sprintf("%s-%d", sn, os.Getpid())),
		ocagent.WithAddress(fmt.Sprintf("%s://%s:%s", ocaSvcPro, ocaSvcAddr, ocaSvcPort)),
	)
	if err != nil {
		return err
	}
	trace.RegisterExporter(oce)
	view.RegisterExporter(oce)
	return err
}
