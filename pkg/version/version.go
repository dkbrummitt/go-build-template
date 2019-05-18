/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package version

import (
	"fmt"
)

// VERSION is the global version string, which should be substituted with a
// real value during build.
var VERSION = "UNKNOWN"

// ReleaseDate is the global release date string, which should be
// substituted with a real value during build.
var RELEASE_DATE = "UNKNOWN"

// GoVersion is the global go lang version string, which should be
// substituted with a real value during build.
var GoVersion = "UNKNOWN"

// GetVersion global func that returns the formatted version string
func GetVersion() string {
	var vFormat = "v%s"
	var rFormat = "Released %s"
	var gFormat = "(Go %s)"
	var spf = "a"

	if RELEASE_DATE != "" && RELEASE_DATE != "UNKNOWN" {
		vFormat = vFormat + " " + rFormat
		spf += "b"
	}
	if GoVersion != "" && GoVersion != "UNKNOWN" {
		vFormat = vFormat + " " + gFormat
		spf += "c"
	}
	switch spf { //build version string based on what is provided
	case "a":
		return fmt.Sprintf(vFormat, VERSION)
	case "ab":
		return fmt.Sprintf(vFormat, VERSION, RELEASE_DATE)
	case "ac":
		return fmt.Sprintf(vFormat, VERSION, GoVersion)
	default:
		return fmt.Sprintf(vFormat, VERSION, RELEASE_DATE, GoVersion)
	}
}

func GetVersionSimple() string {
	var vFormat = "v%s"

	return fmt.Sprintf(vFormat, VERSION)
}
