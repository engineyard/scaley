package steps

import (
	//"fmt"

	"github.com/ess/kennel"
)

type Locking struct{}

func (steps *Locking) StepUp(s kennel.Suite) {
	s.Step(`^a scaling lockfile exists for the group$`, func() error {
		return nil
	})

}

func init() {
	kennel.Register(new(Locking))
}
