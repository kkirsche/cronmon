// Copyright © 2018 Kevin Kirsche <kev.kirsche[at]gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/kkirsche/cronmon/libdb"
	"github.com/kkirsche/cronmon/libmetrics"
	"github.com/kkirsche/cronmon/libserver"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	_ "github.com/lib/pq"
)

var listenPort int

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Starts the cronmon server component",
	Long: `The server command starts the cronmon server component, which runs the
API which receives task connections, can be used to view data, etc.`,
	Run: func(cmd *cobra.Command, args []string) {
		libdb.CreateDBSchema()
		libmetrics.RegisterPrometheusMetrics()

		err := libserver.Run(listenPort)
		if err != nil {
			logrus.WithError(err).Errorln("failed to run server")
		}
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().IntVarP(&listenPort, "listen-port", "p", 8080, "The port for the web server to listen on (overridden by setting the LISTEN_PORT environment variable)")
}
