package steps

import (
	"fmt"

	"github.com/ess/jamaica"
	"github.com/ess/kennel"

	"github.com/engineyard/scaley/pkg/basher"
)

type ScalingDown struct{}

func (steps *ScalingDown) StepUp(s kennel.Suite) {
	s.Step(`^conditions dictate that downscaling is necessary$`, func() error {
		basher.Run = func(string) int {
			return 1
		}

		return nil
	})

	s.Step(`^no messages are logged$`, func() error {
		if len(jamaica.LastCommandOutput()) > 0 {
			fmt.Println("output:", jamaica.LastCommandOutput())
			return fmt.Errorf("A message was logged")
		}

		return nil
	})

}

func init() {
	kennel.Register(new(ScalingDown))
}
