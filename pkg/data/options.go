package data

type Options interface {
	// Rule of thumb, accept interfaces (or bytes) as parameters
	// return struct, bools, (concrete types)
	Validate() (bool, error)
}
