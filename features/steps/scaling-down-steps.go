package steps

import (
	//"fmt"

	"github.com/ess/kennel"
)

type ScalingDown struct{}

func (steps *ScalingDown) StepUp(s kennel.Suite) {
	s.Step(`^conditions dictate that downscaling is necessary$`, func() error {
		return nil
	})

	s.Step(`^there is capacity for the group to downscale$`, func() error {
		return nil
	})

	s.Step(`^the group is scaled down$`, func() error {
		return nil
	})

	s.Step(`^there is not capacity for the group to downscale$`, func() error {
		return nil
	})

	s.Step(`^no messages are logged$`, func() error {
		return nil
	})

}

func init() {
	kennel.Register(new(ScalingDown))
}
