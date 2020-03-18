package data

import "github.com/pkg/errors"

// Config contains settings and configurations support
// various data/storage integrations.
type Config struct {
	Host   string
	DBName string
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
	ok := false
	var err error

	if (Config{}) == c {
		err = errors.New("config not initialized")
		return ok, err
	}
	if c.Host == "" {
		err = errors.New("missing/empty data host")
		return ok, err
	}
	if c.DBName == "" {
		err = errors.New("missing/empty data name")
		return ok, err
	}

	ok = true
	return ok, err
}
