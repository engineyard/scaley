// Copyright © 2018 Engine Yard, Inc.
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
	"fmt"

	"github.com/spf13/cobra"

	"github.com/engineyard/scaley/pkg/workflows"
)

var scaleCmd = &cobra.Command{
	Use: "scale <group>",

	Short: "Scale a group up or down",

	Long: `Scale a group up or down

Given a group description file, writen in YAML, determine if the group in
question should be scaled up or down. If yes, scale the group in the
proper direction.`,

	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("Usage: scaley scale groupname")
		}

		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		return workflows.Perform(
			&workflows.ScalingAGroup{
				GroupName: args[0],
			},
		)
	},

	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	RootCmd.AddCommand(scaleCmd)
}
