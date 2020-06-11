// Licensed to Preferred Networks, Inc. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Preferred Networks, Inc. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// getQueueStateCmd represents the get queue-state command
var getQueueStateCmd = &cobra.Command{
	Use:   "get-queue-state [queue]",
	Short: "Get Queue State",
	Long:  `Get Queue State`,
	Args:  cobra.ExactArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		displayCmdOptsIfEnabled()
		mustInitializeQueueBackend()
	},
	Run: func(cmd *cobra.Command, args []string) {
		queueName := args[0]
		queue, err := queueBackend.GetQueueByName(cmdContext, queueName)
		if err != nil {
			logger.Fatal().Err(err).Msg("Failed to get queue state")
		}
		fmt.Println(queue.Spec.State)
	},
}

func init() {
	rootCmd.AddCommand(getQueueStateCmd)
}