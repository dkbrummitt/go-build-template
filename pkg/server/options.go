package server

import (
	"io/ioutil"

	"github.com/dkbrummitt/go-build-template/pkg/version"
	"github.com/sirupsen/logrus"
)

// Options contains feature switches to support
// various server integrations. This implements
// pkg.Options
type Options struct {
	HasProfiling bool
	HasPush      bool
	CertFile     string
	KeyFile      string
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
} // of LoadDefaults

// DefaultLogger builds a default logrus logger
//
// Pre-Condition:
// - None
// Post-Condition:
// - None
// Params:
// - None
// Returns:
// - logrus entry as a logger
// Errors:
// - None
// Dev Notes:
// - None
func DefaultLogger() *logrus.Entry {
	logger := logrus.New()
	log := logger.WithFields(logrus.Fields{
		"version": version.GetVersionSimple(),
	})

	return log
} // of DefaultLogger
