package error

import (
	"strconv"
)

type AppError struct {
	Code    int
	Message string
	Cause   error
	Action  string
}

type TransientError struct {
	Code       int
	Message    string
	Cause      error
	RetryCount int
}

func buildString(code int, message string, cause error) (e string) {
	if code != 0 {
		e = e + strconv.Itoa(code)
	}
	if message != "" {
		if len(e) > 0 {
			e = e + " "
		}
		e = e + message
	}

	if cause != nil {
		if len(e) > 0 {
			e = e + " "
		}
		e = e + "cause: " + cause.Error()
	}

	return
}

func (ae *AppError) Error() string {
	var e string

	// build error string based provided info
	e = buildString(ae.Code, ae.Message, ae.Cause)

	// add corrective action if provided
	if ae.Action != "" {
		if len(e) > 0 {
			e = e + " "
		}
		e = e + "corrective action: " + ae.Action
	}
	return e
}

func (te *TransientError) Error() string {
	var e string

	// build error string based provided info
	e = buildString(te.Code, te.Message, te.Cause)

	// add corrective action if provided
	if te.RetryCount > 0 {
		if len(e) > 0 {
			e = e + " "
		}
		e = e + "retry count: " + strconv.Itoa(te.RetryCount)
	}
	return e
}
