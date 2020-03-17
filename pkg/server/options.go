package server

import (
	"fmt"
	"io/ioutil"
)

const (
	defaultServerTimeout = 30
)

// Options contains feature switches to support
// various server integrations. This implements
// pkg.Options
type Options struct {
	HasProfiling bool
	HasPush      bool
	CertFile     string
	KeyFile      string
	Timeout      int
} //of Options

// Validate verifys that
// - required option attributes are set
// - attributes that are provided have valid values
//
// Pre-Condition:
// - None
// Post-Condition:
// - None
// Params:
// - None
// Returns:
// - true if valid options provided, otherwise false
// Errors:
// - if one or more errors are detected
// Dev Notes:
// - None
func (o Options) Validate() (bool, error) {
	var ok = false
	var err error

	if (Options{}) == o {
		//nothing set, mark as ok, and quit early
		ok = true
		return ok, err
	}

	// if cert is provided but not found
	if o.CertFile != "" {
		//check for file
		_, err := ioutil.ReadFile(o.CertFile)
		if err != nil {
			//quit early
			return ok, err
		}
	}

	// if key is provided but not found
	if o.KeyFile != "" {
		//check for file
		_, err := ioutil.ReadFile(o.KeyFile)
		if err != nil {
			//quit early
			return ok, err
		}
	}

	ok = true
	return ok, err
} // of Validate

// LoadDefaults loads any default option values
//
// Pre-Condition:
// - Options are initialized
// Post-Condition:
// - Sets HashPush feature flag based on presense and finding cert and key files
// Params:
// - None
// Returns:
// - None
// Errors:
// - None
// Dev Notes:
// - The HasPush feature flag is Read Only. The cert and key files must be valid before it toggles to true
func (o *Options) LoadDefaults() {
	if o == nil {
		//dont do anything
		return
	}
	var err error
	var err2 error
	// if cert is provided but not found
	if o.CertFile != "" {
		//check for file
		_, err = ioutil.ReadFile(o.CertFile)
	}

	// if key is provided but not found
	if o.KeyFile != "" {
		//check for file
		_, err2 = ioutil.ReadFile(o.KeyFile)
	}

	if err == nil && err2 == nil {
		o.HasPush = true
	}

	if o.Timeout == 0 {
		o.Timeout = defaultServerTimeout
	}
} // of LoadDefaults

// String provides a string representation of the state of this struct
//
// Pre-Condition:
// - None
// Post-Condition:
// - None
// Params:
// - None
// Returns:
// - string representation
// Errors:
// - None
// Dev Notes:
// - None
func (o Options) String() string {
	oFmt := "HasProfiling:%v HasPush:%v Timeout:%v"
	return fmt.Sprintf(oFmt, o.HasProfiling, o.HasPush, o.Timeout)
}
