package steps

import (
	"fmt"
	"strings"

	"github.com/ess/jamaica"
	"github.com/ess/kennel"
)

type ScaleySteps struct{}

func (steps *ScaleySteps) StepUp(s kennel.Suite) {
	s.Step(`^I see some output$`, func() error {
		output := jamaica.LastCommandStdout()

		fmt.Println("output:", output)

		return nil
	})

	s.Step(`^I see the help description$`, func() error {
		output := jamaica.LastCommandStdout()

		if !strings.Contains(output, "Generalized faux autoscaling for Engine Yard servers") {
			return fmt.Errorf("Help description not found")
		}

		return nil
	})

	s.Step(`^I see the usage$`, func() error {
		output := jamaica.LastCommandStdout()

		if !strings.Contains(output, "Usage:\n  scaley [command]") {
			return fmt.Errorf("Usage not found")
		}

		return nil
	})

	s.Step(`^I see the available commands$`, func() error {
		output := jamaica.LastCommandStdout()

		if !strings.Contains(output, "Available Commands:\n") {
			return fmt.Errorf("Available commands not found")
		}

		return nil
	})
}

func init() {
	kennel.Register(new(ScaleySteps))
}
