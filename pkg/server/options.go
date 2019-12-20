package server

import (
	"github.com/dkbrummitt/go-build-template/pkg/version"
	"github.com/sirupsen/logrus"
)

// Options contains feature switches to support
// various server integrations. This implements
// pkg.Options
type Options struct {
	Log          *logrus.Entry
	HasProfiling bool
	HasPush      bool
}

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

	ok = true
	return ok, err
}

func defaultLogger() *logrus.Entry {
	logger := logrus.New()
	log := logger.WithFields(logrus.Fields{
		"version": version.GetVersionSimple(),
	})

	return log
}

func (o *Options) LoadDefaults() {
	if o.Log == nil {
		o.Log = defaultLogger()
	}
}
