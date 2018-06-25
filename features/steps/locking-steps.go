package steps

import (
	"fmt"

	"github.com/ess/kennel"
	"github.com/spf13/afero"

	"github.com/engineyard/scaley/pkg/common"
)

type Locking struct{}

func (steps *Locking) StepUp(s kennel.Suite) {
	s.Step(`^a scaling lockfile exists for the group$`, func() error {
		err := common.Root.MkdirAll(common.Locks(), 0755)
		if err != nil {
			return fmt.Errorf("Could not create scaley lock dir")
		}

		return afero.WriteFile(
			common.Root,
			common.Locks()+"/mygroup",
			[]byte(""),
			0644,
		)
	})

	s.Step(`^the group remains locked$`, func() error {
		if !common.FileExists(common.Locks() + "/mygroup") {
			return fmt.Errorf("There is no lockfile for the group")
		}

		return nil
	})

}

func init() {
	kennel.Register(new(Locking))
}
