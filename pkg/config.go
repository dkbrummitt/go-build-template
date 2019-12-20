package pkg

// Config contains settings and configurations support
// various server integrations.
type Config interface {

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
	Validate() (bool, error)
}
