package commands

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/engineyard/scaley/pkg/scaley"
	"github.com/engineyard/scaley/pkg/scaley/bash"
	"github.com/engineyard/scaley/pkg/scaley/eycore"
	"github.com/engineyard/scaley/pkg/scaley/fs"
	"github.com/engineyard/scaley/pkg/scaley/http"
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

		reportingURL = viper.GetString("reporting_url")
		if len(reportingURL) == 0 {
			return fmt.Errorf(
				`You must provide a reporting URL.
				
This should be listed as reporting_url: in /etc/scaley/config.yml`,
			)
		}

		token = viper.GetString("token")
		if len(token) == 0 {
			return fmt.Errorf(
				`This operation requires Engine Yard API authentication.
				
This should be listed as token: in /etc/scaley/config.yml`,
			)
		}

		api = viper.GetString("api")
		if api == "" {
			api = "https://api.engineyard.com"
		}

		return nil
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		return scaley.Finalize(
			scaley.Scale().Call(
				&scaley.ScalingEvent{
					GroupName: args[0],
					Services:  services(),
				},
			),
		)
	},

	SilenceUsage:  true,
	SilenceErrors: true,
}

func init() {
	RootCmd.AddCommand(scaleCmd)
}

var services = func() *scaley.Services {
	eycore.Setup(api, token)

	return &scaley.Services{
		Groups:       fs.NewGroupService(),
		Servers:      eycore.NewServerService(),
		Environments: eycore.NewEnvironmentService(),
		Scripts:      fs.NewScalingScriptService(),
		Locker:       fs.NewLockService(),
		Runner:       bash.NewExecService(),
		Log:          http.NewLogService(reportingURL),
	}
}

var api string
var reportingURL string
var token string

// Copyright © 2019 Engine Yard, Inc.
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
