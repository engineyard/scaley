package steps

import (
	"fmt"

	"github.com/ess/kennel"
	"github.com/spf13/afero"

	"github.com/engineyard/scaley/v2/pkg/scaley/fs"
)

type Locking struct{}

func (steps *Locking) StepUp(s kennel.Suite) {
	s.Step(`^a scaling lockfile exists for the group$`, func() error {
		err := fs.Root.MkdirAll(fs.Locks(), 0755)
		if err != nil {
			return fmt.Errorf("Could not create scaley lock dir")
		}

		return afero.WriteFile(
			fs.Root,
			fs.Locks()+"/mygroup",
			[]byte(""),
			0644,
		)
	})

	s.Step(`^the group remains locked$`, func() error {
		if !fs.FileExists(fs.Locks() + "/mygroup") {
			return fmt.Errorf("There is no lockfile for the group")
		}

		return nil
	})

	s.Step(`^the group is unlocked$`, func() error {
		if fs.FileExists(fs.Locks() + "/mygroup") {
			return fmt.Errorf("There is a lockfile for the group")
		}

		return nil
	})

}

func init() {
	kennel.Register(new(Locking))
}
