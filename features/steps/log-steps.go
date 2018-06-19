package steps

import (
	"fmt"
	"strings"

	"github.com/ess/jamaica"
	"github.com/ess/kennel"
	"github.com/ess/mockable"
)

type LogSteps struct{}

func (steps *LogSteps) StepUp(s kennel.Suite) {
	s.Step(`^a scaling failure is logged$`, func() error {
		if !strings.Contains(jamaica.LastCommandStdout(), "FAILURE : Group[mygroup]: Could not be scaled up - Errors occurred while starting these servers, please contact support: ") {
			return fmt.Errorf("Warning not found")
		}

		return nil
	})

	s.Step(`^a chef failure is logged$`, func() error {
		if !strings.Contains(jamaica.LastCommandStdout(), "FAILURE : Group[mygroup]: Could not be scaled up - A Chef error occurred while upscaling the group. Please contact support.") {
			return fmt.Errorf("Chef failure not found")
		}

		return nil
	})

	s.BeforeSuite(func() {
		mockable.Enable()
	})

	s.AfterSuite(func() {
		mockable.Disable()
	})

}

func init() {
	kennel.Register(new(LogSteps))
}
