package strategies

import (
	"strings"

	"github.com/engineyard/scaley/pkg/scaley"
)

type legion struct {
	group scaley.Group
	ops   scaley.OpsService
}

func newLegion(group scaley.Group, ops scaley.OpsService) *legion {
	return &legion{group: group, ops: ops}
}

func (scaler *legion) Upscale() error {
	failures := make([]string, 0)

	for _, candidate := range newBallot(scaler.group).All(scaley.Up) {
		err := scaler.ops.Start(candidate)
		if err != nil {
			failures = append(failures, candidate.ProvisionedID)
		}
	}

	if len(failures) > 0 {
		return fmt.Errorf(upFail, strings.Join(failures, ", "))
	}

	return nil
}

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
