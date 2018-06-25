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

package util

import (
	"fmt"

	"github.com/spf13/viper"
)

func Token() string {
	return viper.GetString("token")
}

func ApiHost() string {
	api := viper.GetString("api")
	if api == "" {
		return "api.engineyard.com"
	}

	return api
}

func CheckToken() error {
	if len(Token()) == 0 {
		return fmt.Errorf(
			`This operation requires Engine Yard API authentication.

This should be listed as token: in /etc/scaley/config.yml`,
		)
	}

	return nil
}

func ReportingURL() string {
	return viper.GetString("reporting_url")
}

func CheckReportingURL() error {
	if len(ReportingURL()) == 0 {
		return fmt.Errorf(
			`You must provide a reporting URL.

This should be listed as reporting_url: in /etc/scaley/config.yml`,
		)
	}

	return nil
}
