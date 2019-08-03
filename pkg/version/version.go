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
func GetVersion() string {
	var vFormat = "%s"
	var rFormat = "Released %s"
	var gFormat = "(Go %s)"
	var spf = "a"

	if RELEASE_DATE != "" && RELEASE_DATE != "UNKNOWN" {
		vFormat = vFormat + " " + rFormat
		spf += "b"
	}
	if GO_VERSION != "" && GO_VERSION != "UNKNOWN" {
		vFormat = vFormat + " " + gFormat
		spf += "c"
	}
	switch spf { // build version string based on what is provided
	case "a":
		return fmt.Sprintf(vFormat, VERSION)
	case "ab":
		return fmt.Sprintf(vFormat, VERSION, RELEASE_DATE)
	case "ac":
		return fmt.Sprintf(vFormat, VERSION, GO_VERSION)
	default:
		return fmt.Sprintf(vFormat, VERSION, RELEASE_DATE, GO_VERSION)
	}
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
	return VERSION
}
