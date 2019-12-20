package version

import (
	"fmt"
	"net/http"
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
		router.Handle("/version", http.HandlerFunc(handler))
		router.Handle("/version/", http.HandlerFunc(handler))
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
		"version": "%s",
		"go": "%s",
		"releasDate": "%s",
		"commit": "%s"
	}`
	p := fmt.Sprintf(f, VERSION, GO_VERSION, RELEASE_DATE, GIT_COMMIT)
	w.Write([]byte(p))
}
