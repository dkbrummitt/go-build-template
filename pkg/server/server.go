package server

import (
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
}

// Run Starts the server with routes based on the provided configurations and options
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
		cert, key, err = loadCerts(s.Config)
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
}

func New(opts Options, conf Config) (s Server, e error) {
	if _, err := conf.Validate(); err != nil {
		return
	}
	if _, err := opts.Validate(); err != nil {
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
}

func loadCerts(c Config) (cert, key string, err error) {
	certB, err := ioutil.ReadFile(c.CertFile)
	if err != nil {
		return //quit early
	}
	keyB, err := ioutil.ReadFile(c.KeyFile)
	if err != nil {
		return //quit early
	}

	cert = string(certB)
	key = string(keyB)

	return
}
