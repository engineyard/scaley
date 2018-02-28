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
	candidates := scaler.group.Candidates("up")

	if len(candidates) == 0 {
		return fmt.Errorf("There are no servers in the group avaiable for upscaling")
	}

	candidate := candidates[0]

	if err := startServer(candidate, scaler.api); err != nil {
		return fmt.Errorf(upFail, candidate.AmazonID())
	}

	return nil
}

func (scaler *individual) Downscale() error {
	candidates := scaler.group.Candidates("down")

	if len(candidates) == 0 {
		return fmt.Errorf("There are no servers in the group avaiable for downscaling")
	}

	candidate := candidates[0]

	if err := stopServer(candidate, scaler.api); err != nil {
		return fmt.Errorf(downFail, candidate.AmazonID())
	}

	return nil
}
