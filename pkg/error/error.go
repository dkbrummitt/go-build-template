package error

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"gopkg.in/yaml.v2"
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

type DefinedErrors struct {
	// Name   string `yaml:name`
	// Detail Detail `yaml:detail`
	Errors map[string]Detail `yaml:"errors,omitempty"`
}

type Detail struct {
	Code      int    `yaml:"code"`
	Action    string `yaml:"action,omitempty"`
	Transient bool   `yaml:"transient"`
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

func (de *DefinedErrors) Load() error {
	var err error

	//load only once
	if len(de.Errors) > 0 {
		//quit early,
		return err
	}

	// read in the error yaml
	f, err := ioutil.ReadFile("errors.yml")
	if err != nil {
		fmt.Println("warning: no defined errors detected")
		return err
	}
	fmt.Println("YAML FILE:\n" + string(f))
	// if successful, marshall errors into Defined errors
	me := make(map[string]Detail)
	err = yaml.Unmarshal(f, &me)
	if err != nil {
		fmt.Println("warning: defined errors detected, but unreadable")
		return err
	}
	fmt.Printf("YAML Struct:\n%+v", me)
	de.Errors = me
	return err
}
