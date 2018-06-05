package main

import (
	"os"
	"testing"
	"time"

	"github.com/engineyard/scaley/features/steps"

	"github.com/DATA-DOG/godog"
	"github.com/engineyard/scaley/cmd/scaley/cmd"
	"github.com/ess/jamaica"
	"github.com/ess/kennel"
	"github.com/ess/mockable"
)

var commandOutput string
var lastCommandRanErr error

func TestMain(m *testing.M) {
	mockable.Enable()
	steps.Register()
	jamaica.SetRootCmd(cmd.RootCmd)

	status := godog.RunWithOptions(
		"godog",
		func(s *godog.Suite) {
			jamaica.StepUp(s)
			kennel.StepUp(s)
		},

		godog.Options{
			Format:    "pretty",
			Paths:     []string{"features"},
			Randomize: time.Now().UTC().UnixNano(),
		},
	)

	if st := m.Run(); st > status {
		status = st
	}

	os.Exit(status)
}

func TestTrue(t *testing.T) {
}
