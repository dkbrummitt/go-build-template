package data

type Config interface {
	// Rule of thumb, accept interfaces (or bytes) as parameters
	// return struct, bools, (concrete types)
	Validate() (bool, error)
}
