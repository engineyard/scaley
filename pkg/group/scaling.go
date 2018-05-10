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
		return fmt.Errorf("A Chef error occurred while upscaling the group. Please contact support.")
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

//func (g *Group) unusableServers() []scaler.Server {
//unusable := make([]scaler.Server, 0)

//for _, s := range g.ScalingServers {
//state := s.Instance.State
//if state != "running" && state != "stopped" {
//unusable = append(unusable, s)
//}
//}

//return unusable
//}

func (g *Group) ScalingStrategy() string {
	return g.Strategy
}
