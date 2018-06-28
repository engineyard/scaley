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

package scaler

import (
	"fmt"
	"strings"

	"github.com/engineyard/eycore/core"
)

type legion struct {
	group scalable
	api   core.Client
}

func newLegion(group scalable, api core.Client) *legion {
	return &legion{group: group, api: api}
}

func (scaler *legion) Upscale() error {
	// 1. Start all scaling servers
	failures := make([]string, 0)

	for _, s := range scaler.group.Candidates("up") {
		err := startServer(s, scaler.api)
		if err != nil {
			failures = append(failures, s.AmazonID())
		}
	}

	if len(failures) > 0 {
		return fmt.Errorf(upFail, strings.Join(failures, ", "))
	}

	return nil
}

func (scaler *legion) Downscale() error {
	failures := make([]string, 0)

	for _, s := range scaler.group.Candidates("down") {
		err := stopServer(s, scaler.api, scaler.group.PreStop())
		if err != nil {
			failures = append(failures, s.AmazonID())
		}
	}

	if len(failures) > 0 {
		return fmt.Errorf(downFail, strings.Join(failures, ", "))
	}

	return nil
}
