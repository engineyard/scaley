package scaler

import (
	"fmt"

	"github.com/engineyard/eycore/core"
)

type single struct {
	group scalable
	api   core.Client
}

func newSingle(group scalable, api core.Client) *single {
	return &single{group: group, api: api}
}

func (scaler *single) Upscale() error {
	candidate := scaler.group.Candidates("up")[0]

	if err := startServer(candidate, scaler.api); err != nil {
		return fmt.Errorf(upFail, candidate.AmazonID())
	}

	return nil
}

func (scaler *single) Downscale() error {
	candidate := scaler.group.Candidates("down")[0]

	if err := stopServer(candidate, scaler.api); err != nil {
		return fmt.Errorf(downFail, candidate.AmazonID())
	}

	return nil
}
