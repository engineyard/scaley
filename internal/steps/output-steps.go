package steps

import (
	"fmt"
	"strings"

	"github.com/ess/jamaica"
	"github.com/ess/kennel"
)

type OutputSteps struct{}

func (steps *OutputSteps) StepUp(s kennel.Suite) {
	s.Step(`^I see an error about the group's missing scaling script$`, func() error {
		output := fmt.Sprintf("%s", jamaica.LastCommandStatus())

		if !strings.Contains(output, "mygroup requires a scaling_script") {
			return fmt.Errorf("Didn't see a scaling script error")
		}

		return nil
	})

	s.Step(`^I see an error about the non-existent scaling script$`, func() error {
		output := fmt.Sprintf("%s", jamaica.LastCommandStatus())

		if !strings.Contains(output, "mygroup scaling_script does not exist on the file system") {
			return fmt.Errorf("Didn't see a scaling script error")
		}

		return nil
	})

	s.Step(`^I see an error about the missing scaling servers$`, func() error {
		output := fmt.Sprintf("%s", jamaica.LastCommandStatus())

		if !strings.Contains(output, "mygroup contains no scaling servers") {
			return fmt.Errorf("Didn't see a scaling server error")
		}

		return nil
	})

	s.Step(`^I see an error about the invalid scaling server$`, func() error {
		output := fmt.Sprintf("%s", jamaica.LastCommandStatus())

		if !strings.Contains(output, "mygroup contains invalid scaling server") {
			return fmt.Errorf("Didn't see a scaling server error")
		}

		return nil
	})

	s.Step(`^the tests print the output$`, func() error {
		fmt.Printf(
			"output: %s\nerror: %s\n",
			jamaica.LastCommandStdout(),
			jamaica.LastCommandStatus(),
		)

		return nil
	})

}

func init() {
	kennel.Register(new(OutputSteps))
}
