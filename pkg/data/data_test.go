package data

import (
	"fmt"
	"strings"
	"testing"
)

func Test_DataError(t *testing.T) {
	tstCases := []struct {
		e string
		c int
		a string
	}{
		{},          // 0
		{"", 0, ""}, // 0

		{"", 0, "fix it"},   // 3
		{"", 123, ""},       // 2
		{"", 123, "fix it"}, // 5

		{"something bad happened", 0, ""},         // 1
		{"something bad happened", 0, "fix it"},   // 4
		{"something bad happened", 123, ""},       // 3
		{"something bad happened", 123, "fix it"}, // 6
	}

	for _, tstCase := range tstCases {
		de := DataError{tstCase.e, tstCase.c, tstCase.a}
		err := de.Error()
		if err == "" && (tstCase.e != "" || tstCase.c > 0 || tstCase.a != "") {
			t.Errorf("Expected err message, but saw ''. Test Case: %+v", tstCase)
		}
		if tstCase.a == "" && strings.Contains(err, "CORRECTIVE ACTION") {
			t.Errorf("Expected err message to not have a corrective action but saw '%s'. Test Case: %+v", err, tstCase)
		}
		if tstCase.a != "" && !strings.Contains(err, "CORRECTIVE ACTION") {
			if !strings.Contains(err, tstCase.a) {
				t.Errorf("Expected err message to have corrective action but saw '%s'. Test Case: %+v", err, tstCase)
			}
			if (tstCase.c != 0 || tstCase.e != "") && !strings.Contains(err, "CORRECTIVE ACTION") {
				t.Errorf("Expected err message to specify corrective action but saw '%s'. Test Case: %+v", err, tstCase)
			}
		}

		if tstCase.c != 0 && !strings.Contains(err, fmt.Sprintf("%d", tstCase.c)) {
			t.Errorf("Expected err message to have a error code but saw '%s'. Test Case: %+v", err, tstCase)
		}
		if tstCase.e != "" && !strings.Contains(err, tstCase.e) {
			t.Errorf("Expected err message to have error message but saw '%s'. Test Case: %+v", err, tstCase)
		}

		if tstCase.e != "" && !strings.Contains(err, tstCase.e) {
			t.Errorf("Expected err message to have error message but saw '%s'. Test Case: %+v", err, tstCase)
		}
	}
}

func Benchmark_DataError(b *testing.B) {
	for ndx := 0; ndx < b.N; ndx++ {
		de := DataError{"something bad happened", 999, "fix it"}
		_ = de.Error()
	}
}
