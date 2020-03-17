package data

import (
	"context"

	"github.com/dkbrummitt/go-build-template/pkg/stats"
	"github.com/sirupsen/logrus"
)

type Data interface {
	// Rule of thumb, accept interfaces (or bytes) as parameters
	// return struct, bools, (concrete types)
	Create(ctx context.Context, data interface{}, log *logrus.Entry) ([]byte, error)
	Retrieve(ctx context.Context, data interface{}, log *logrus.Entry) ([]byte, error)
	List(ctx context.Context, data interface{}, log *logrus.Entry) ([]byte, error)
	Update(ctx context.Context, data interface{}, log *logrus.Entry) error
	Delete(ctx context.Context, data interface{}, log *logrus.Entry) error
	Find(ctx context.Context, filter map[string]interface{}, log *logrus.Entry) ([]byte, error)
	Ping(ctx context.Context, stat *stats.Stats, log *logrus.Entry) error
}

/*
// Sample Implementation of the interface
type dataImpl struct {}

(d *dataImpl) Create(ctx context.Context, data interface{}, log *logrus.Entry) ([]byte, error){
	// place imagination here ...
}

(d *dataImpl) Retrieve(ctx context.Context, data interface{}, log *logrus.Entry) ([]byte, error){
	// place imagination here ...
}

(d *dataImpl) List(ctx context.Context, data interface{}, log *logrus.Entry) ([]byte, error){
	// place imagination here ...
}

(d *dataImpl) Update(ctx context.Context, data interface{}, log *logrus.Entry) error{
	// place imagination here ...
}

(d *dataImpl) Delete(ctx context.Context, data interface{}, log *logrus.Entry) error{
	// place imagination here ...
}

(d *dataImpl) Find(ctx context.Context, filter map[string]interface{}, log *logrus.Entry) ([]byte, error){
	// place imagination here ...
}

(d *dataImpl) Ping(ctx context.Context, stat *stats.Stats, log *logrus.Entry) error{
	// place imagination here ...
}


func NewData() Data {
  return &dataImpl
}
*/
