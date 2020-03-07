package server

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	_ "net/http/pprof"
	"sync"

	"github.com/dkbrummitt/go-build-template/pkg/version"
)

// Server provides http(s) support to this application.
// Set options for feature flags and Configs for env/system
// configurations
type Server struct {
	Options
	Config

	router *http.ServeMux
} //of Server

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

	// load certs if needed
	if s.HasPush {
		cert, key, err = loadCerts(s.Options)
		if err != nil {
			return
		}
	}
	//register handlers/route
	if s.HasProfiling {
		fmt.Println("Profiling enabled")
		s.RegisterHandler("/debug/", http.DefaultServeMux)
	}
	version.RegisterRoute(s.router)
	srv = &http.Server{
		Addr:    port,
		Handler: s.router,
	}

	go func() {
		defer func() {
			fmt.Println("Server Shutdown complete!")
			wg.Done() // notify shutdown is done
		}()

		//start the server
		if s.HasPush {
			err = srv.ListenAndServeTLS(cert, key)
		} else {
			err = srv.ListenAndServe()
		}
		if err != http.ErrServerClosed {
			// unexpected error
			fmt.Println("Run(): ", err)
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
		return
	}
	if ok, err := opts.Validate(); ok && err != nil {
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
func loadCerts(o Options) (cert, key string, err error) {
	certB, err2 := ioutil.ReadFile(o.CertFile)
	if err != nil {
		err = err2
	}
	keyB, err3 := ioutil.ReadFile(o.KeyFile)
	if err3 != nil {
		//maintain all error data
		if err == nil {
			err = err3
		} else {
			err = errors.New(err.Error() + "; " + err3.Error())
		}
	}

	cert = string(certB)
	key = string(keyB)

	return
} // of loadCerts
