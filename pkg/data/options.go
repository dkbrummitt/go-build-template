package data

type Options interface {
	// Rule of thumb, accept interfaces (or bytes)
	// return structs*
	Validate() (bool, error)
}

/*
//
type dataImpl struct {
	Create() (interface{}, error){
		// Place imagination here
	}
	Retrieve() (interface{}, error){
		// Place imagination here
	}
	Update() (interface{}, error){
		// Place imagination here
	}
	Delete() (interface{}, error){
		// Place imagination here
	}
}
func NewData() Data {
  return &dataImpl
}
*/
