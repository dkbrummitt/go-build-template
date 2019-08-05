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
		v string // version
	}{
		// NOTE truth-table like pattern to organiz test conditions.
		// Not required, but it does help to orgnize thoughts
		{""},
		{"UNKNOWN"},
		{"0.0.0"},
	}

	for _, tstCase := range tstCases {
		VERSION = tstCase.v

		r := GetVersionSimple()
		e := tstCase.v

		if tstCase.v != "" && tstCase.v != "UNKNOWN" && r != e {
			t.Errorf("Expected version to be '%s'. Saw instead '%s'.  Case: '%s'", e, r, tstCase.v)
		}
		if (tstCase.v == "" || tstCase.v == "UNKNOWN") && r != "NO-STATIC-VERSION" {
			t.Errorf("Expected version to be 'NO-STATIC-VERSION'. Saw instead '%s'.  Case: '%s' ", r, tstCase.v)
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
		gc string
	}{
		// NOTE truth-table like pattern to organiz test conditions.
		// Not required, but it does help to orgnize thoughts
		{"UNKNOWN", "UNKNOWN", "UNKNOWN", "UNKNOWN"}, // default nothing is set via ld-flags
		{"UNKNOWN", "UNKNOWN", "UNKNOWN", ""},
		{"UNKNOWN", "UNKNOWN", "", "UNKNOWN"},
		{"UNKNOWN", "UNKNOWN", "", ""},
		{"UNKNOWN", "", "UNKNOWN", "UNKNOWN"},
		{"UNKNOWN", "", "UNKNOWN", ""},
		{"UNKNOWN", "", "", "UNKNOWN"},
		{"UNKNOWN", "", "", ""},

		{"", "UNKNOWN", "UNKNOWN", "UNKNOWN"},
		{"", "UNKNOWN", "UNKNOWN", ""},
		{"", "UNKNOWN", "", "UNKNOWN"},
		{"", "UNKNOWN", "", ""},
		{"", "", "UNKNOWN", "UNKNOWN"},
		{"", "", "UNKNOWN", ""},
		{"", "", "", "UNKNOWN"},
		{"", "", "", ""},

		{"0.0.0", "2019-02-30", "1.12", "ab12cd34"},
		{"0.0.0", "2019-02-30", "1.12", ""},
		{"0.0.0", "2019-02-30", "", "ab12cd34"},
		{"0.0.0", "2019-02-30", "", ""},
		{"0.0.0", "", "1.12", "ab12cd34"},
		{"0.0.0", "", "1.12", ""},
		{"0.0.0", "", "", "ab12cd34"},
		{"0.0.0", "", "", ""},

		{"", "2019-02-30", "1.12", "ab12cd34"},
		{"", "2019-02-30", "1.12", ""},
		{"", "2019-02-30", "", "ab12cd34"},
		{"", "2019-02-30", "", ""},
		{"", "", "1.12", "ab12cd34"},
		{"", "", "1.12", ""},
		{"", "", "", "ab12cd34"},
		{"", "", "", ""},
	}

	for _, tstCase := range tstCases {
		VERSION = tstCase.v
		RELEASE_DATE = tstCase.rd
		GO_VERSION = tstCase.gv
		GIT_COMMIT = tstCase.gc

		r := GetVersion()
		// e := tstCase.v

		if r == "" {
			t.Errorf("Expected version significant. Saw instead ''. Case: '%s' '%s' '%s'", tstCase.v, tstCase.rd, tstCase.gv)
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
