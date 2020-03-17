/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
package cmd

import (
	"encoding/json"
	"fmt"
	"time"

	"dkbrummitt/go-build-template/pkg/stats"
	"github.com/spf13/cobra"
)

// statsCmd represents the stats command
var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Snapshot of OS health",
	Long: `Snapshot of OS Memory and Disk, with additional info.

	THIS COMMAND IS FOR DEMO PURPOSES ONLY
	IT IS NOT INTENDED FOR PRODUCTION.`,
	Run: func(cmd *cobra.Command, args []string) {

		st := stats.NewStats(stats.Options{
			StartTime: time.Now(),
		})
		st.PullStats()
		// b, err := json.Marshal(st)
		b, err := json.MarshalIndent(st, "", "  ")
		if err != nil {
			fmt.Println("Error marshalling stats:", err)
		}
		json.MarshalIndent(st, "", "  ")
		fmt.Println(string(b))
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)

	// Here you will define your flags and configuration settings.
}
