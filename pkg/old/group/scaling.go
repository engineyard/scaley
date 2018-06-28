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

package group

import (
	"fmt"

	"github.com/engineyard/eycore/core"
	"github.com/engineyard/eycore/environments"

	"github.com/engineyard/scaley/pkg/common"
	"github.com/engineyard/scaley/pkg/scaler"
)

func (group *Group) CanScale(direction string) bool {
	switch direction {
	case "up":
		if len(group.candidatesForUpscale()) > 0 {
			return true
		}
	case "down":
		if len(group.candidatesForDownscale()) > 0 {
			return true
		}
	}

	return false
}

func Scale(group *Group, api core.Client, direction string) error {
	if direction == "up" {
		return upscale(group, api)
	}

	return downscale(group, api)
}

func upscale(group *Group, api core.Client) error {
	// 1. Scale up with the group's defined strategy (default: legion)
	err := scaler.For(group, api).Upscale()
	if err != nil {
		return err
	}

	// 2. Run chef on the environment
	err = runChef(api, group.Environment)
	if err != nil {
		return fmt.Errorf("A Chef error occurred while upscaling the group. Please contact support.")
	}

	return nil
}

func runChef(api core.Client, environment *environments.Model) error {
	req, err := environments.Apply(api, environment, "")
	if err != nil {
		return err
	}

	req, err = common.WaitFor(api, req)
	if err != nil {
		return err
	}

	if !req.Successful {
		return fmt.Errorf("%s", req.RequestStatus)
	}

	return nil
}

func downscale(group *Group, api core.Client) error {
	// 1. Scale down with the group's defined strategy (default: legion)
	err := scaler.For(group, api).Downscale()
	if err != nil {
		return err
	}

	// 2. Run chef on the environment
	err = runChef(api, group.Environment)
	if err != nil {
		return fmt.Errorf("A Chef error occurred while downscaling the group. Please contact support.")
	}

	return nil
}

func (g *Group) Candidates(direction string) []scaler.Server {
	if direction == "up" {
		return g.candidatesForUpscale()
	}

	return g.candidatesForDownscale()
}

func (g *Group) candidatesForUpscale() []scaler.Server {
	candidates := make([]scaler.Server, 0)

	for _, s := range g.ScalingServers {
		if s.Instance.State == "stopped" {
			candidates = append(candidates, s)
		}
	}

	return candidates
}

func (g *Group) candidatesForDownscale() []scaler.Server {
	candidates := make([]scaler.Server, 0)

	for _, s := range g.ScalingServers {
		if s.Instance.State == "running" {
			candidates = append(candidates, s)
		}
	}

	return candidates
}

func (g *Group) ScalingStrategy() string {
	return g.Strategy
}

func (g *Group) PreStop() string {
	return g.StopScript
}
