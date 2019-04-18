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

package main

import (
	"encoding/json"
	"fmt"

	"github.com/dkbrummitt/go-build-template/pkg/stats"
	"github.com/dkbrummitt/go-build-template/pkg/version"
)

func main() {
	//sample of version data that can be logged or sent to monitoring/usage-tracking
	fmt.Println("MyApp", version.GetVersion())

	//Sample pull of stats data that can be logged or sent to monitoring
	fmt.Println("System Stats:")
	st := stats.NewStats(stats.StatsOptions{})
	st.PullStats()
	// b, err := json.Marshal(st)
	b, err := json.MarshalIndent(st, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling stats:", err)
	}
	json.MarshalIndent(st, "", "  ")
	fmt.Println(string(b))
}
