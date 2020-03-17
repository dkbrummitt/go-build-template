package version

import (
	"fmt"
	"net/http"
)

var (
	path = "/api/v1/version"
)

// RegisterRoute Add server route support for paths '/version' and '/version/
//
// Pre-Condition:
// - router is not nil
// Post-Condition:
// - server will support /version with and without the trailing path
// Params:
// - router http.ServerMux
// Returns:
// - None
// Errors:
// - None
// Dev Notes:
// - This func is a simple setter and deemed too-small-to-test
func RegisterRoute(router *http.ServeMux) {
	if router != nil {
		// support with and without trailing slash
		router.Handle(path, http.HandlerFunc(handler))
		router.Handle(path+"/", http.HandlerFunc(handler))
	}
}

// handler builds the response for the version route/path
//
// Pre-Condition:
// - func has been added to a http.ServeMux
// Post-Condition:
// - None
// Params:
// - router http.ServerMux
// Returns:
// - None
// Errors:
// - None
// Dev Notes:
// - This func is a simple string builder and deemed too-small-to-test
func handler(w http.ResponseWriter, r *http.Request) {
	f := `{
		"name": "%s",
		"version": "%s",
		"go": "%s",
		"releasDate": "%s",
		"commit": "%s"
	}`

	// w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Accept")
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case "OPTIONS":
		//nothing to do, headers already set
	case "GET":
		p := fmt.Sprintf(f, AppName, VERSION, GO_VERSION, RELEASE_DATE, GIT_COMMIT)
		w.Write([]byte(p))
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
