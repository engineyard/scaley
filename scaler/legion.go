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
		return fmt.Errorf("Errors occurred while starting these servers, please contact support: %s", strings.Join(failures, ", "))
	}

	return nil
}

func (scaler *legion) Downscale() error {
	failures := make([]string, 0)

	for _, s := range scaler.group.Candidates("down") {
		err := stopServer(s, scaler.api)
		if err != nil {
			failures = append(failures, s.AmazonID())
		}
	}

	if len(failures) > 0 {
		return fmt.Errorf("Errors occurred while stopping these servers, please contact support: %s", strings.Join(failures, ", "))
	}

	return nil
}
