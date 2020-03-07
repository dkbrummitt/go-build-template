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

	"github.com/dkbrummitt/go-build-template/pkg/server"
	"github.com/spf13/cobra"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start application HTTP(S) Server",
	Long:  `Provide support as a server instead of ad-hoc/one-off requests`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		fmt.Println("serve called")
		o := server.Options{}
		c := server.Config{}
		c.Port, err = cmd.Flags().GetInt("port")
		o.HasProfiling, err = cmd.Flags().GetBool("profile")

		s, err := server.NewServer(o, c)
		if err != nil {
			fmt.Println(err)
			return
		}

		httpServerExitDone := &sync.WaitGroup{}
		httpServerExitDone.Add(1)
		srvr, err := s.Run(httpServerExitDone)
		if err != nil {
			fmt.Println(err)
			return
		}

		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
		go func() {
			for range ch {
				fmt.Println("sudo shutdown now!")
				if err := srvr.Shutdown(context.TODO()); err != nil {
					panic(err) // failure/timeout shutting down the server gracefully
				}

				// wait for server to stop
				fmt.Printf("main: done. exiting")
			}
		}()

		httpServerExitDone.Wait() // wait for it all to end

	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().IntP("port", "p", 8080, "Port Number for server. Default 8000")
	serveCmd.Flags().BoolP("profile", "r", false, "Enable profiling. Default false")
}
