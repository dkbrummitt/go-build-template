package version

import (
	"fmt"
	"testing"
)

// Test_GetVersionSimple_NoSet  test for performance testing to
// determine potential improvements for this function
func Test_GetVersionSimple(t *testing.T) {
	// conditions where function will succeed
	tbls := []struct {
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

	for _, tbl := range tbls {
		VERSION = tbl.v
		RELEASE_DATE = tbl.rd
		GO_VERSION = tbl.gv

		r := GetVersionSimple()
		e := tbl.v

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
	tbls := []struct {
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

	if false { // For Compile sake.  Remove after address FIXMEs
		fmt.Printf("Test arguments %+v", tbls)
	}
	// FIX ME Is GetVersion Testable? If so, write the test for it below
	// FIX ME If not, write the test for it below
}

// Benchmark_GetVersion Benchmark test for performance testing to
// determine potential improvements for this function
func Benchmark_GetVersion(b *testing.B) {
	// make a bunch of calls
	for i := 0; i < b.N; i++ {
		GetVersion()
	}
}
