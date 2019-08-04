package main

import (
	"encoding/json"
	"fmt"

	"github.com/dkbrummitt/go-build-template/pkg/stats"
	"github.com/dkbrummitt/go-build-template/pkg/version"
)

func main() {
	// sample of version data that can be logged or sent to monitoring/usage-tracking
	fmt.Println("MyApp", version.GetVersion())

	// Sample pull of stats data that can be logged or sent to monitoring
	fmt.Println("System Stats:")
	st := stats.NewStats(stats.Options{})
	st.PullStats()
	// b, err := json.Marshal(st)
	b, err := json.MarshalIndent(st, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling stats:", err)
	}
	json.MarshalIndent(st, "", "  ")
	fmt.Println(string(b))
}
