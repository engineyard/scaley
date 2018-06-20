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
		found := 0
		output := jamaica.LastCommandStdout()

		if strings.Contains(output, "FAILURE : Group[mygroup]: Could not be scaled up - Errors occurred while starting these servers, please contact support: ") {
			found += 1
		}

		if strings.Contains(output, "FAILURE : Group[mygroup]: Could not be scaled down - Errors occurred while stopping these servers, please contact support: ") {
			found += 1
		}

		if found == 0 {
			return fmt.Errorf("Failure not found")
		}

		return nil
	})

	s.Step(`^a chef failure is logged$`, func() error {
		found := 0
		output := jamaica.LastCommandStdout()

		if strings.Contains(output, "FAILURE : Group[mygroup]: Could not be scaled up - A Chef error occurred while upscaling the group. Please contact support.") {
			found += 1
		}

		if strings.Contains(output, "FAILURE : Group[mygroup]: Could not be scaled down - A Chef error occurred while downscaling the group. Please contact support.") {
			found += 1
		}

		if found == 0 {
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
