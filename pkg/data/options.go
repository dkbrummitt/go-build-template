package data

// Options contains feature switches to support
// various data/storage integrations.
type Options struct {
	Timeout int
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
	ok := false
	var err error

	return ok, err
}
