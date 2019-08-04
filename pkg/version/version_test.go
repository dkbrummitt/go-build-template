package version

import (
	"testing"
)

// Some tests/benchmarks in this file are too small to test. They are
// included to provide samples of how to
// - write table driven tests
// - write benchmarks

// Test_GetVersionSimple_NoSet  test for performance testing to
// determine potential improvements for this function.
// GetVersionSimple is considered too small to test.
func Test_GetVersionSimple(t *testing.T) {
	// conditions where function will succeed
	tstCases := []struct {
		v  string // version
		rd string // release(build) date
		gv string // go version
	}{
		// NOTE truth-table like pattern to organiz test conditions.
		// Not required, but it does help to orgnize thoughts
		{"UNKNOWN", "UNKNOWN", "UNKNOWN"}, // default nothing is set via ld-flags
		{"UNKNOWN", "UNKNOWN", ""},
		{"UNKNOWN", "", "UNKNOWN"},
		{"UNKNOWN", "", ""},
		{"", "UNKNOWN", "UNKNOWN"},
		{"", "UNKNOWN", ""},
		{"", "", "UNKNOWN"},
		{"", "", ""},

		{"UNKNOWN", "UNKNOWN", "1.12"},
		{"UNKNOWN", "Feb 30, 2020", "UNKNOWN"},
		{"UNKNOWN", "Feb 30, 2020", "1.12"},
		{"1.0.0", "UNKNOWN", "UNKNOWN"},
		{"1.0.0", "UNKNOWN", "1.12"},
		{"1.0.0", "Feb 30, 2020", "UNKNOWN"},
		{"1.0.0", "Feb 30, 2020", "1.12"},
	}

	for _, tstCase := range tstCases {
		VERSION = tstCase.v
		RELEASE_DATE = tstCase.rd
		GO_VERSION = tstCase.gv

		r := GetVersionSimple()
		e := tstCase.v

		if r != e {
			t.Errorf("Expected version to be '%s'. Saw instead '%s'", e, r)
		}
	}

}

// Benchmark_GetVersionSimple Benchmark test for performance testing to
// determine potential improvements for this function
func Benchmark_GetVersionSimple(b *testing.B) {
	// make a bunch of calls
	for i := 0; i < b.N; i++ {
		GetVersionSimple()
	}
}

// Test_GetVersion Benchmark test for performance testing to
// determine potential improvements for this function
func Test_GetVersion(t *testing.T) {
	// conditions where function will succeed
	tstCases := []struct {
		v  string
		rd string
		gv string
	}{
		// NOTE truth-table like pattern to organiz test conditions.
		// Not required, but it does help to orgnize thoughts
		{"UNKNOWN", "UNKNOWN", "UNKNOWN"}, // default nothing is set via ld-flags
		{"UNKNOWN", "UNKNOWN", ""},
		{"UNKNOWN", "", "UNKNOWN"},
		{"UNKNOWN", "", ""},
		{"", "UNKNOWN", "UNKNOWN"},
		{"", "UNKNOWN", ""},
		{"", "", "UNKNOWN"},
		{"", "", ""},

		{"UNKNOWN", "UNKNOWN", "1.12"},
		{"UNKNOWN", "Feb 30, 2020", "UNKNOWN"},
		{"UNKNOWN", "Feb 30, 2020", "1.12"},
		{"1.0.0", "UNKNOWN", "UNKNOWN"},
		{"1.0.0", "UNKNOWN", "1.12"},
		{"1.0.0", "Feb 30, 2020", "UNKNOWN"},
		{"1.0.0", "Feb 30, 2020", "1.12"},
	}

	for _, tstCase := range tstCases {
		VERSION = tstCase.v
		RELEASE_DATE = tstCase.rd
		GO_VERSION = tstCase.gv

		r := GetVersion()
		// e := tstCase.v

		if r == "" {
			t.Error("Expected version significant. Saw instead ''")
		}
	}
}

// Benchmark_GetVersion Benchmark test for performance testing to
// determine potential improvements for this function
func Benchmark_GetVersion(b *testing.B) {
	// make a bunch of calls
	for i := 0; i < b.N; i++ {
		GetVersion()
	}
}
