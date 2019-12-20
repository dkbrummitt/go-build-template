package error

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func Test_AppError(t *testing.T) {
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
		var tmpError error
		de := AppError{tstCase.cde, tstCase.msg, tstCase.coz, tstCase.act}
		err := de.Error()
		tmpError = &de
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

		//err should be a AppError
		switch e := tmpError.(type) {
		case nil:
			t.Errorf("Expected err, got nil test case: %+v", tstCase)
		case *AppError:
			// its all good in the hood
			fmt.Println("We're good!")
		default:
			t.Errorf("Expected DataErr, got somehthing different: %s %+v", e, tstCase)
		}
		// fmt.Println(fmt.Sprintf("Error '%s'", err))
	}
}

func Benchmark_AppError(b *testing.B) {
	rc := errors.New("i am the glitch in the matrix")
	for ndx := 0; ndx < b.N; ndx++ {
		de := AppError{999, "something bad happened", rc, "fix it"}
		_ = de.Error()
	}
}

func Test_TransientError(t *testing.T) {
	rc := errors.New("i am the glitch in the matrix")
	tstCases := []struct {
		msg string
		cde int
		r   int
		coz error
	}{
		{}, // empty error

		{"", 0, 0, nil}, // empty error
		{"", 0, 0, rc},

		{"", 0, 1, nil},
		{"", 0, 1, rc},

		{"", 1, 0, nil},
		{"", 1, 0, rc},

		{"", 1, 1, nil},
		{"", 1, 1, rc},

		{"something bad happened", 0, 0, nil},
		{"something bad happened", 0, 0, rc},

		{"something bad happened", 0, 1, nil},
		{"something bad happened", 0, 1, rc},

		{"something bad happened", 1, 0, nil},
		{"something bad happened", 1, 0, rc},

		{"something bad happened", 1, 1, nil},
		{"something bad happened", 1, 1, rc},
	}

	for _, tstCase := range tstCases {
		var tmpError error
		de := TransientError{tstCase.cde, tstCase.msg, tstCase.coz, tstCase.r}
		err := de.Error()
		tmpError = &de
		// err message is empty, but err is significant
		if err == "" && (tstCase.msg != "" || tstCase.cde > 0 || tstCase.r != 0 || tstCase.coz != nil) {
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

		if tstCase.r != 0 && !strings.Contains(err, strconv.Itoa(tstCase.r)) {
			t.Errorf("Expected err message to have corrective action but saw '%s'. Test Case: %+v", err, tstCase)
		}

		if tstCase.coz != nil && !strings.Contains(err, tstCase.coz.Error()) {
			t.Errorf("Expected err message to contain root cause but saw '%s'. Test Case: %+v", err, tstCase)
		}

		//all double spaces are removed
		if strings.Contains(err, "  ") {
			t.Errorf("Expected err message to not contain any double spaces '%s'. Test Case: %+v", err, tstCase)
		}

		//err should be a AppError
		switch e := tmpError.(type) {
		case nil:
			t.Errorf("Expected err, got nil test case: %+v", tstCase)
		case *TransientError:
			// its all good in the hood
			fmt.Println("We're good!")
		default:
			t.Errorf("Expected DataErr, got somehthing different: %s %+v", e, tstCase)
		}
	}
}

func Benchmark_TransientError(b *testing.B) {
	rc := errors.New("i am the glitch in the matrix")
	for ndx := 0; ndx < b.N; ndx++ {
		de := TransientError{999, "something bad happened", rc, ndx}
		_ = de.Error()
	}
}
