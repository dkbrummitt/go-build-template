package data

import (
	"errors"
	"strconv"
	"strings"
	"testing"
)

func Test_DataError(t *testing.T) {
	rc := errors.New("i am the glitch in the matrix")
	tstCases := []struct {
		msg string
		cde int
		act string
		coz error
	}{
		{}, // empty error

		{"", 0, "", nil}, // empty error
		{"", 0, "", rc},

		{"", 0, "fix it", nil},
		{"", 0, "fix it", rc},

		{"", 1, "", nil},
		{"", 1, "", rc},

		{"", 1, "fix it", nil},
		{"", 1, "fix it", rc},

		{"something bad happened", 0, "", nil},
		{"something bad happened", 0, "", rc},

		{"something bad happened", 0, "fix it", nil},
		{"something bad happened", 0, "fix it", rc},

		{"something bad happened", 1, "", nil},
		{"something bad happened", 1, "", rc},

		{"something bad happened", 1, "fix it", nil},
		{"something bad happened", 1, "fix it", rc},
	}

	for _, tstCase := range tstCases {
		de := DataError{tstCase.cde, tstCase.msg, tstCase.act, tstCase.coz}
		err := de.Error()

		// err message is empty, but err is significant
		if err == "" && (tstCase.msg != "" || tstCase.cde > 0 || tstCase.act != "" || tstCase.coz != nil) {
			t.Errorf("Expected err message, but saw ''. Test Case: %+v", tstCase)
		}

		// error message should only contain significant codes (non-zero)
		if tstCase.cde == 0 && strings.HasPrefix(err, "0") {
			t.Errorf("Expected err message not contain error code but saw '%s'. Test Case: %+v", err, tstCase)
		}

		// expect error message to contain any significant data provided
		if tstCase.cde != 0 && !strings.HasPrefix(err, strconv.Itoa(tstCase.cde)) {
			// if code is significant, then it should be in the error message
			t.Errorf("Expected err message to have an error code but saw '%s'. Test Case: %+v", err, tstCase)
		}

		if tstCase.msg != "" && !strings.Contains(err, tstCase.msg) {
			t.Errorf("Expected err message to have error message but saw '%s'. Test Case: %+v", err, tstCase)
		}

		if tstCase.act != "" && !strings.Contains(err, tstCase.act) {
			t.Errorf("Expected err message to have corrective action but saw '%s'. Test Case: %+v", err, tstCase)
		}

		if tstCase.coz != nil && !strings.Contains(err, tstCase.coz.Error()) {
			t.Errorf("Expected err message to contain root cause but saw '%s'. Test Case: %+v", err, tstCase)
		}

		//all double spaces are removed
		if strings.Contains(err, "  ") {
			t.Errorf("Expected err message to not contain any double spaces '%s'. Test Case: %+v", err, tstCase)
		}

		// fmt.Println(fmt.Sprintf("Error '%s'", err))
	}
}

func Benchmark_DataError2(b *testing.B) {
	rc := errors.New("i am the glitch in the matrix")
	for ndx := 0; ndx < b.N; ndx++ {
		de := DataError{999, "something bad happened", "fix it", rc}
		_ = de.Error()
	}
}
