package server

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// Config contains settings and configurations support
// various server integrations. Implements
// pkg.Config
type Config struct {
	Port   int
	Log    *logrus.Entry
	Logger *logrus.Logger
}

// Validate verifys that
// - required config attributes are set
// - attributes that are provided have valid values
//
// Pre-Condition:
// - None
// Post-Condition:
// - None
// Params:
// - None
// Returns:
// - true if valid configs provided, otherwise false
// Errors:
// - if Config is not initialized
// - if one or more errors are detected
// Dev Notes:
// - None
func (c Config) Validate() (bool, error) {
	var ok = false
	var err error

	if (Config{}) == c {
		err = errors.New("server configuration not initialized")
		return ok, err
	}

	if !c.validPort() {
		eFmt := "server configuration port provided(%d) is invalid"
		err = errors.New(fmt.Sprintf(eFmt, c.Port))
		return ok, err
	}

	if c.Log == nil {
		err = errors.New("context log not provided")
		return ok, err
	}

	if c.Logger == nil {
		err = errors.New("logger not provided")
		return ok, err
	}

	// checks passed, set ok to true
	ok = true
	return ok, err
}

// validPort verifies the port provided is within a port range. Note
// a valid port is provided when it is between the 1025 and 65535, inclusive.
// this allows for both registered ports (1025-49151)
// and dynamic ports (49152-65535)
//
// Pre-Condition:
// - None
// Post-Condition:
// - None
// Params:
// - None
// Returns:
// - true if valid port provided, otherwise false
// Errors:
// - None
// Dev Notes:
// - None
func (c Config) validPort() (ok bool) {
	if c.Port >= 1025 && c.Port <= 65535 {
		ok = true
	}

	return
}

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
func (c Config) String() string {
	cFmt := ":%d"
	return fmt.Sprintf(cFmt, c.Port)
}
