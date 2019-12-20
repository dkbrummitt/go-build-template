package pkg

// Options contains feature switches to support
// various server integrations.
type Options interface {

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
	Validate() (bool, error)
}
