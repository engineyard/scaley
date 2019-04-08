package steps

import (
	"fmt"
	"strings"

	"github.com/ess/jamaica"
	"github.com/ess/kennel"
	"github.com/ess/mockable"
)

type ScalingUp struct{}

func (steps *ScalingUp) StepUp(s kennel.Suite) {
	s.Step(`^conditions dictate that upscaling is necessary$`, func() error {
		stubBasher(2)

		return nil
	})

	s.Step(`^a warning is logged regarding the insufficient capacity$`, func() error {
		if !strings.Contains(jamaica.LastCommandStdout(), "WARNING : Group[mygroup]: Cannot be scaled up - Consider adding more servers to the group") {
			return fmt.Errorf("Warning not found")
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
	kennel.Register(new(ScalingUp))
}
