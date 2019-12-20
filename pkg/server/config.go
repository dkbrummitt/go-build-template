package server

import (
	"fmt"
	"io/ioutil"
)

// Config contains settings and configurations support
// various server integrations. Implements
// pkg.Config
type Config struct {
	CertFile string
	KeyFile  string
	Port     int
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
// - if one or more errors are detected
// Dev Notes:
// - None
func (c Config) Validate() (bool, error) {
	var ok = false
	var err error

	// if cert is provided but not found
	if c.CertFile != "" {
		//check for file
		_, err := ioutil.ReadFile(c.CertFile)
		if err != nil {
			//quit early
			return ok, err
		}
	}

	// if key is provided but not found
	if c.KeyFile != "" {
		//check for file
		_, err := ioutil.ReadFile(c.KeyFile)
		if err != nil {
			//quit early
			return ok, err
		}
	}

	if !validPort(c.Port) {
		err = fmt.Errorf("Port provided(%d) is invalid", c.Port)
		return ok, err
	}
	ok = true
	return ok, err
}

func (c *Config) LoadDefaults() {

}

func validPort(port int) (ok bool) {
	if port >= 1 && port <= 65535 {
		ok = true
	}

	return
}
