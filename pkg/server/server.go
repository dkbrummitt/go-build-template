package server

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"

	"github.com/dkbrummitt/go-build-template/pkg/version"
	"github.com/pkg/errors"
)

// Server provides http(s) support to this application.
// Set options for feature flags and Configs for env/system
// configurations
type Server struct {
	Options
	Config

	router *http.ServeMux
} //of Server

func DefaultHeaders(w http.ResponseWriter) {
	allowedHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, Authorization, X-CSRF-Token, If-Modified-Since, If-Unmodified-Since, If-None-Match"

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, GET, OPTIONS, PATCH, POST, PUT")
	w.Header().Set("Access-Control-Allow-Headers", allowedHeaders)
	w.Header().Set("Server", "") // clear the server
}

// RegisterHandler allow clients to register new routes/paths
//
// Pre-Condition:
// - Server is initialized
// - Server has valid options
// - Server has valid configs
// Post-Condition:
// - router has new path with associated Handler
// Params:
// - path string subpath for this route
// - handler Handler that will apply logic for request and associated response
// Returns:
// - None
// Errors:
// - None
// Dev Notes:
// - None
func (s *Server) RegisterHandler(path string, handler http.Handler) {
	s.router.Handle(path, handler)
} // of RegisterHandler

// Run Starts the server with routes based on the provided configurations and
// options. If cert and key info is provided in options, the server will start
// as an TLS (Push) server.
//
// Pre-Condition:
// - Server is initialized
// - Server has valid options
// - Server has valid configs
// Post-Condition:
// - server is running and able to accept requests
// Params:
// - None
// Returns:
// - Error
// Errors:
// - if certs/TLS is indicated but not found
// - if errors while starting TLS server support
// - if errors while starting non-TLS server support
// Dev Notes:
// - None
func (s Server) Run(wg *sync.WaitGroup) (srv *http.Server, err error) {
	port := fmt.Sprintf(":%d", s.Port)
	var cert string
	var key string
	w := s.Logger.Writer()
	defer w.Close()

	s.Log.Infof("starting server:%+s", s.String())

	// load certs if needed
	if s.HasPush {
		cert, key, err = s.Options.loadCerts()
		if err != nil {
			err = errors.Wrap(err, "load cert files failed")
			return
		}
	}
	//register handlers/route
	if s.HasProfiling {
		s.Log.Warn("Profiling enabled")
		s.RegisterHandler("/debug/", http.DefaultServeMux)
	}
	version.RegisterRoute(s.router)
	srv = &http.Server{
		ErrorLog:     log.New(w, "", 0),
		Addr:         port,
		Handler:      s.router,
		ReadTimeout:  time.Duration(s.Options.Timeout) * time.Second,
		WriteTimeout: time.Duration(s.Options.Timeout) * time.Second,
	}

	go func() {
		defer func() {
			s.Log.Warn("Server Shutdown complete!")
			wg.Done() // notify shutdown is done
		}()

		//start the server
		if s.HasPush {
			err = srv.ListenAndServeTLS(cert, key)
			err = errors.Wrap(err, "start of tls server failed")
		} else {
			err = srv.ListenAndServe()
			err = errors.Wrap(err, "start of nontls server failed")
		}
		if err != http.ErrServerClosed {
			// unexpected error
			s.Log.Error("error starting server", err)
		}
	}()

	return
} // of Run

// NewServer initializes a new server based on options and configs provided
//
// Pre-Condition:
// - None
// Post-Condition:
// - None
// Params:
// - opts Options or feature flags supported for this server
// - conf Config or required settings for this server
// Returns:
// - instance of server
// - errors if any
// Errors:
// - invalid port provided
// Dev Notes:
// - None
func NewServer(opts Options, conf Config) (s Server, e error) {
	if ok, err := conf.Validate(); ok && err != nil {
		err = errors.Wrap(err, "server config validation failed")
		return
	}
	if ok, err := opts.Validate(); ok && err != nil {
		err = errors.Wrap(err, "server options validation failed")
		return
	}
	s = Server{
		Options: opts,
		Config:  conf,
	}

	// load defaults
	s.Options.LoadDefaults()

	//init the router
	s.router = http.NewServeMux()

	return
} // of NewServer

// loadCerts loads/reads certificate files provided
// Pre-Condition:
// - options param provided is properly initalized
// Post-Condition:
// - None
// Params:
// - Options, the server options
// Returns:
// - The servers certificate file if found
// - The servers key file if found
// Errors:
// - when failure reading cert file
// - when failure reading key file
// Dev Notes:
// - None
func (o Options) loadCerts() (cert, key string, err error) {
	certB, err2 := ioutil.ReadFile(o.CertFile)
	if err != nil {
		err = err2
		err = errors.Wrap(err2, "error reading certificate file in configuration")
	}
	keyB, err3 := ioutil.ReadFile(o.KeyFile)
	if err3 != nil {
		//maintain all error data
		if err == nil {
			err = err3
			err = errors.Wrap(err3, "error reading key file in configuration")
		} else {
			err = errors.New(err.Error() + "; " + err3.Error())
		}
	}

	cert = string(certB)
	key = string(keyB)

	return
} // of loadCerts

// String provides a string representation of the state of this struct
//
// Pre-Condition:
// - None
// Post-Condition:
// - None
// Params:
// - None
// Returns:
// - string representation
// Errors:
// - None
// Dev Notes:
// - None
func (s Server) String() string {
	sFmt := "Config:%v Options:%v"
	return fmt.Sprintf(sFmt, s.Config.String(), s.Options.String())
}
