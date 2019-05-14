package metrics

import (
	"context"
	"fmt"
	"os"
	"strconv"
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

func InitMetrics(sn string, tc trace.Config) {
	_, err := initJaegerTracing(sn)
	if err != nil {
		fmt.Println("WARNING unable to initialize Jaeger Tracing", err.Error())
		//clear the error
		err = nil
	}
	initTraceConfigs(tc)

	rptPeriod := 0
	if os.Getenv("METRICS_REPORTING_PERIOD") != "" {
		rptPeriod, err = strconv.Atoi(os.Getenv("METRICS_REPORTING_PERIOD"))
		if err != nil {
			fmt.Println("WARNING env variable METRICS_REPORTING_PERIOD is not a number:", os.Getenv("METRICS_REPORTING_PERIOD"))
		}
	}
	initViewReporting(rptPeriod)
}

//initTraceConfigs Defaults to 10% of requests... recommended tracing should be closer to 1%
//per google folks. By default open census samples 1 in 10000 traces.
//this function overrides that default. NOTE if you are just running locally, then AlwaysSample is the default
func initTraceConfigs(tc trace.Config) {
	if env == "local" && tc.DefaultSampler == nil {
		trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	}

	if tc.DefaultSampler != nil {
		trace.ApplyConfig(tc)
	} else {
		//set default config
		trace.ApplyConfig(trace.Config{DefaultSampler: trace.ProbabilitySampler(defaultProbability)})
	}
} //initTraceConfigs

func initViewReporting(rptPeriod int) {

	view.SetReportingPeriod(defaultReportingPeriod * time.Second)
	if rptPeriod != 0 {
		view.SetReportingPeriod(time.Duration(rptPeriod) * time.Second)
	}
} //initViewReporting

func AddLatencyView(name string, tags ...tag.Key) (*view.View, error) {
	latencyView := &view.View{
		Name:        name + "/latency",
		Measure:     getMeasure(name),
		Description: "The distribution of the latencies",

		// Latency in buckets:
		// [>=0ms, >=25ms, >=50ms, >=75ms, >=100ms, >=200ms, >=400ms, >=600ms, >=800ms, >=1s, >=2s, >=4s, >=6s]
		Aggregation: view.Distribution(0, 25, 50, 75, 100, 200, 400, 600, 800, 1000, 2000, 4000, 6000),
		TagKeys:     tags}
	err := view.Register(latencyView)

	return latencyView, err
} //AddLatencyView

func getMeasure(name string) *stats.Float64Measure {
	latencyMS := stats.Float64(fmt.Sprintf("%s/latency", name), "The latency in milliseconds per "+name+" loop", "ms")
	return latencyMS
} //getMeasure

func sinceInMilliseconds(startTime time.Time) float64 {
	return float64(time.Since(startTime).Nanoseconds()) / 1e6
} //sinceInMilliseconds

/* Code sample for spanning/tracing with stats
 */
func doSomeSpanning(ctx context.Context) error {
	startTime := time.Now()
	funcName := "doSomeSpanning"
	var err error
	if ctx == nil {
		ctx = context.Background()
	}

	//Sample usage in client files

	ctx, s := trace.StartSpan(ctx, funcName)
	//record how long it took
	defer stats.Record(ctx, getMeasure(funcName).M(sinceInMilliseconds(startTime)))
	defer s.End()
	//set tags
	ctx, err = tag.New(ctx, tag.Insert(KeyFunction, funcName))

	//place imagination here...

	if err != nil {
		s.SetStatus(trace.Status{
			Code:    trace.StatusCodeUnknown,
			Message: err.Error(),
		})
		ctx, err = tag.New(ctx, tag.Upsert(KeyError, err.Error()), tag.Insert(KeyStatus, "ERROR"))
		return err
	}
	s.SetStatus(trace.Status{
		Code: trace.StatusCodeOK,
	})
	ctx, err = tag.New(ctx, tag.Upsert(KeyStatus, "OK"))
	return err
}

/* // */
