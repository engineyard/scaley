package strategies

import (
	"github.com/engineyard/scaley/pkg/scaley"
)

type individual struct {
	group scaley.Group
	ops   scaley.OpsService
}

func newIndividual(group scaley.Group, ops scaley.OpsService) *individual {
	return &individual{group: group, ops: ops}
}

func (scaler *individual) Upscale() error {
	candidate := newBallot(scaler.group).Single(scaley.Up)

	if err := scaler.ops.Start(candidate); err != nil {
		return fmt.Errorf(upFail, candidate.ProvisionedID)
	}

	return nil
}

func (scaler *individual) Downscale() error {
	candidate := newBallot(scaler.group).Single(scaley.Down)

	if err := scaler.ops.Stop(candidate); err != nil {
		return fmt.Errorf(downFail, candidate.ProvisionedID)
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
