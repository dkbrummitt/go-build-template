package stats

import (
	"fmt"
	"testing"
	"time"
)

func Test_NewStats(t *testing.T) {
	now := time.Now()
	cases := []struct {
		st time.Time //start time
	}{
		{},
		{now},
	}

	for _, tstCase := range cases {
		opts := Options{
			StartTime: tstCase.st,
		}
		result := NewStats(opts)

		//If starttime isnt set, then I should not have app details to start
		if tstCase.st.IsZero() && result.AppDetails != nil {
			t.Errorf("Expected no App Details, got %+v", result.AppDetails)
		}

		// if starttime is set, then I should have app details
		if !tstCase.st.IsZero() && result.AppDetails == nil {
			t.Errorf("Expected App Details, got nil")
		}

		// if starttime is set, and app details is set, then I should have
		// startime set in appdetails
		if !tstCase.st.IsZero() && result.AppDetails != nil {
			if _, ok := result.AppDetails["startTime"]; !ok {
				t.Errorf("Expected AppDetails.startTime to be set, got instead %+v", result.AppDetails)
			}
		}
	}
}

// Benchmark_NewStats perf test creating new stats. Creating options is
// treated as incidental, so is not included in the benchmark loop
func Benchmark_NewStats(b *testing.B) {

	opts := Options{
		StartTime: time.Now(),
	}
	for i := 0; i < b.N; i++ {
		NewStats(opts)
	}
}

func Test_PullStats(t *testing.T) {

	opts := Options{
		StartTime: time.Now(),
	}
	tst1 := NewStats(opts)
	err := tst1.PullStats()
	if err != nil {
		t.Errorf("Expected nil error, instead got,%s", err)
	}
	if _, ok := tst1.AppDetails["uptime"]; !ok {
		t.Errorf("Expected AppDetails.upstime to be set, instead saw nil: %+v", tst1)
	}
	if up, ok := tst1.AppDetails["uptime"]; ok && up == "UNKNOWN" {
		t.Error("Expected AppDetails.upstime to be significant, instead saw UNKNOWN")
	}
	fmt.Printf("STATS: %+v", &tst1)
}

func Benchmark_PullStats(b *testing.B) {
	opts := Options{
		StartTime: time.Now(),
	}
	tst := NewStats(opts)
	for i := 0; i < b.N; i++ {
		tst.PullStats()
	}
}
