package data

import (
	"context"

	"github.com/dkbrummitt/go-build-template/pkg/stats"
)

type Data interface {
	// Rule of thumb, accept interfaces (or bytes) as parameters
	// return struct, bools, (concrete types)
	Create(context.Context, []byte) error
	Retrieve(context.Context, interface{}) (interface{}, error)
	Update(context.Context, []byte, interface{}) error
	Delete(context.Context, interface{}) error
	Find(context.Context, map[string]interface{}) (interface{}, error)
	Ping(context.Context, *stats.Stats) error
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
