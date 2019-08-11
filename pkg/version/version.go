package version

import (
	"fmt"
)

var (
	// VERSION is the global version string, which should be substituted with a
	// real value during build via LDFLAGS.
	VERSION = "UNKNOWN"

	// RELEASE_DATE is the global release date string, which should be
	// substituted with a real value during build via LDFLAGS.
	RELEASE_DATE = "UNKNOWN"

	// GO_VERSION is the global go lang version string, which should be
	// substituted with a real value during build via LDFLAGS.
	GO_VERSION = "UNKNOWN"

	// GIT_COMMIT it the git scm commit sha that was compiled. Th
	GIT_COMMIT = "UNKNOWN"
)

// GetVersion global func that returns the formatted detailed version string
// that contains version, release date, and go version used for this build.
// Possible formats include:
//  1.0.2
//  1.0.2 2019-03-31
//	1.02 Go 1.12.7
//	1.02 2019-03-31 Go 1.12.7
//
// Pre-Condition:
// - None
// Post-Condition:
// - None
// Params:
// - None
// Returns:
// - Formatted string representation of the version, that includes
//   version number, release/build date, and go lang version.
// Errors:
// - None
// Dev Notes:
// - None
func GetVersion() (v string) {

	if VERSION == "" || VERSION == "UNKNOWN" {
		VERSION = "NO-STATIC-VERSION"
	}
	v = VERSION

	if GIT_COMMIT != "" && GIT_COMMIT != "UNKNOWN" {
		gcFormat := "Commit %s"
		v = v + " " + gcFormat
		v = fmt.Sprintf(v, GIT_COMMIT)
	}
	if RELEASE_DATE != "" && RELEASE_DATE != "UNKNOWN" {
		rFormat := "Released %s"
		v = v + " " + rFormat
		v = fmt.Sprintf(v, RELEASE_DATE)
	}
	if GO_VERSION != "" && GO_VERSION != "UNKNOWN" {
		gFormat := "(Go %s)"
		v = v + " " + gFormat
		v = fmt.Sprintf(v, GO_VERSION)
	}
	return
}

// GetVersionSimple global func that returns the formatted detailed version
// string that contains version.
//
// Pre-Condition:
// - None
// Post-Condition:
// - None
// Params:
// - None
// Returns:
// - Formatted string representation of the version, that includes
//   only version number. This is intended for use for metrics and heatlh.
// Errors:
// - None
// Dev Notes:
//	- Having this as a function is overkill... I should probably nix it
func GetVersionSimple() string {

	if VERSION == "" || VERSION == "UNKNOWN" {
		VERSION = "NO-STATIC-VERSION"
	}

	return VERSION
}
