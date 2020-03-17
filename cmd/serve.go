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
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"dkbrummitt/go-build-template/pkg/logs"
	"dkbrummitt/go-build-template/pkg/server"

	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start application HTTP(S) Server",
	Long:  `Provide support as a server instead of ad-hoc/one-off requests`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		srvOpts := server.Options{}
		srvConf := server.Config{}

		asJSON, _ := cmd.Flags().GetBool("json")
		reportCaller, _ := cmd.Flags().GetBool("caller")
		level, _ := cmd.Flags().GetString("level")

		srvConf.Port, _ = cmd.Flags().GetInt("port")
		srvOpts.HasProfiling, _ = cmd.Flags().GetBool("profile")
		srvConf.Log, srvConf.Logger = logs.NewLogger(asJSON, reportCaller, level)

		srv, err := server.NewServer(srvOpts, srvConf)
		if err != nil {
			srvConf.Log.Error("unable to create server with error ", err)
			return
		}

		httpServerExitDone := &sync.WaitGroup{}
		httpServerExitDone.Add(1)
		srvr, err := srv.Run(httpServerExitDone)
		if err != nil {

			srvConf.Log.Error("unsable to start server with error ", err)
			return
		}

		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
		go func() {
			for range ch {
				fmt.Println("sudo shutdown now!")
				if err := srvr.Shutdown(context.TODO()); err != nil {
					srvConf.Log.Warnf("timeout hit during shutdown with error %s", err)
					panic(err) // failure/timeout shutting down the server gracefully
				}

				// wait for server to stop
				srvConf.Log.Warn("main: done. exiting")
			}
		}()

		httpServerExitDone.Wait() // wait for it all to end

	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Server
	serveCmd.Flags().IntP("port", "p", 8080, "Port Number for server. Default 8000")
	serveCmd.Flags().BoolP("profile", "r", false, "Enable profiling. Default false")
	serveCmd.Flags().IntP("timeout", "t", 30, "Timeout for server. Default 30 secs")

	// Logging
	serveCmd.Flags().BoolP("json", "j", true, "Enable JSON format logging. Default true")
	serveCmd.Flags().BoolP("caller", "c", false, "Enable report caller logging. CAUSES SLOW PERFORMANCE Default false")
	serveCmd.Flags().StringP("level", "l", "debug", "Specify log level. Default debug")
}
