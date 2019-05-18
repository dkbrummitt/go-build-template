package metrics

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
	"go.opencensus.io/trace"
)

const (
	defaultProtocol        = "http"
	defaultReportingPeriod = 60 //in seconds
	defaultProbability     = .10
)

var (
	env = os.Getenv("ENV")
	//DefaultKeys/Tags
	KeyApp, _      = tag.NewKey("application")
	KeyVersion, _  = tag.NewKey("version")
	KeyRegion, _   = tag.NewKey("region")
	KeyHost, _     = tag.NewKey("host")
	KeyIP, _       = tag.NewKey("ip")
	KeyMethod, _   = tag.NewKey("method")
	KeyFunction, _ = tag.NewKey("function")
	KeyStatus, _   = tag.NewKey("status")
	KeyError, _    = tag.NewKey("error")
)

type MetricOptions struct {
	ServiceName string
	Config      trace.Config
	Mux         *http.ServeMux
}

func InitMetrics(opts MetricOptions) {
	fmt.Println("Initializing telemetry collection")
	initTraceConfigs(opts)
	var err error
	// err := initPrometheusTracing(opts)
	// if err != nil {
	// 	fmt.Println("WARNING unable to initialize Prometheus", err.Error())
	// 	//clear the error
	// 	err = nil
	// }

	err = initJaegerTracing(opts)
	if err != nil {
		fmt.Println("WARNING unable to initialize Jaeger", err.Error())
		//clear the error
		err = nil
	}

	// rptPeriod := 0
	// if os.Getenv("METRICS_REPORTING_PERIOD") != "" {
	// 	rptPeriod, err = strconv.Atoi(os.Getenv("METRICS_REPORTING_PERIOD"))
	// 	if err != nil {
	// 		fmt.Println("WARNING env variable METRICS_REPORTING_PERIOD is not a number:", os.Getenv("METRICS_REPORTING_PERIOD"))
	// 	}
	// }
	// initViewReporting(rptPeriod)
}

//initTraceConfigs Defaults to 10% of requests... recommended tracing should be closer to 1%
//per google folks. By default open census samples 1 in 10000 traces.
//this function overrides that default. NOTE if you are just running locally, then AlwaysSample is the default
func initTraceConfigs(opts MetricOptions) {
	fmt.Println("Initializing trace configurations...")
	if env == "local" && opts.Config.DefaultSampler == nil {
		fmt.Println("\talways sample")
		trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	}

	if opts.Config.DefaultSampler != nil {
		fmt.Println("\tprovided sample")
		trace.ApplyConfig(opts.Config)
	} else {
		//set default config
		fmt.Println("\tprobablity sample")
		trace.ApplyConfig(trace.Config{DefaultSampler: trace.ProbabilitySampler(defaultProbability)})
	}
} //initTraceConfigs

func initViewReporting(rptPeriod int) {
	fmt.Println("Initializing view reporting...")
	view.SetReportingPeriod(defaultReportingPeriod * time.Second)
	if rptPeriod != 0 {
		view.SetReportingPeriod(time.Duration(rptPeriod) * time.Second)
	}
} //initViewReporting

func AddLatencyView(name string, tags ...tag.Key) (*view.View, error) {
	fmt.Println("Adding latency view...")
	latencyView := &view.View{
		Name:        name + "/latency",
		Measure:     GetMeasure(name),
		Description: "The distribution of the latencies",

		// Latency in buckets:
		// [>=0ms, >=25ms, >=50ms, >=75ms, >=100ms, >=200ms, >=400ms, >=600ms, >=800ms, >=1s, >=2s, >=4s, >=6s]
		Aggregation: view.Distribution(0, 25, 50, 75, 100, 200, 400, 600, 800, 1000, 2000, 4000, 6000),
		TagKeys:     tags}
	err := view.Register(latencyView)

	return latencyView, err
} //AddLatencyView

func GetMeasure(name string) *stats.Float64Measure {

	latencyMS := stats.Float64(fmt.Sprintf("%s/latency", name), "The latency in milliseconds per "+name+" loop", "ms")
	return latencyMS
} //GetMeasure

func SinceInMilliseconds(startTime time.Time) float64 {
	return float64(time.Since(startTime).Nanoseconds()) / 1e6
} //SinceInMilliseconds
