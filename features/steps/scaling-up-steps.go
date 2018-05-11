package steps

import (
	//"fmt"

	"github.com/ess/kennel"
)

type ScalingUp struct{}

func (steps *ScalingUp) StepUp(s kennel.Suite) {
	s.Step(`^conditions dictate that upscaling is necessary$`, func() error {
		return nil
	})

	s.Step(`^a warning is logged regarding the insufficient capacity$`, func() error {
		return nil
	})

}

func init() {
	kennel.Register(new(ScalingUp))
}
