package steps

import (
	"fmt"
	"strings"

	"github.com/ess/jamaica"
	"github.com/ess/kennel"
	//"github.com/engineyard/scaley/pkg/basher"
)

type DownScript struct{}

func (steps *DownScript) StepUp(s kennel.Suite) {

	s.Step(`^the stop script is not executed$`, func() error {
		if strings.Contains(jamaica.LastCommandOutput(), "STOP_SCRIPT:") {
			return fmt.Errorf("Expected the stop script to not be run")
		}

		return nil
	})

	s.Step(`^the stop script is executed for each target server$`, func() error {
		output := jamaica.LastCommandOutput()

		expected := len(mygroup.ScalingServers)
		actual := strings.Count(output, "STOP_SCRIPT")

		if actual != expected {
			return fmt.Errorf("Expected %d stop script results, got %d", expected, actual)
		}

		return nil
	})

	s.Step(`^a stop script failure is logged for the first server$`, func() error {
		output := jamaica.LastCommandOutput()

		expected := 1
		actual := strings.Count(output, "STOP_SCRIPT_ERROR")

		if actual != expected {
			return fmt.Errorf("Expected %d stop script errors, got %d", expected, actual)
		}

		found := ""

		for _, line := range strings.Split(output, "\n") {
			if strings.HasPrefix(line, "FAILURE : Group[mygroup]:") {
				found = line
			}
		}

		if len(found) == 0 {
			return fmt.Errorf("No failure generated")
		}

		if !strings.Contains(found, mygroup.ScalingServers[0].ID) {
			return fmt.Errorf("First server not listed in failure")
		}

		return nil
	})

}

func init() {
	kennel.Register(new(DownScript))
}
