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

	"github.com/engineyard/eycore/core"
)

type individual struct {
	group scalable
	api   core.Client
}

func newIndividual(group scalable, api core.Client) *individual {
	return &individual{group: group, api: api}
}

func (scaler *individual) Upscale() error {
	candidate := scaler.group.Candidates("up")[0]

	if err := startServer(candidate, scaler.api); err != nil {
		return fmt.Errorf(upFail, candidate.AmazonID())
	}

	return nil
}

func (scaler *individual) Downscale() error {
	candidate := scaler.group.Candidates("down")[0]

	if err := stopServer(candidate, scaler.api, scaler.group.PreStop()); err != nil {
		return fmt.Errorf(downFail, candidate.AmazonID())
	}

	return nil
}
