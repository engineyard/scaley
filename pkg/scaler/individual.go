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
