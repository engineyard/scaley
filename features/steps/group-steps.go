package steps

import (
	//"fmt"

	"github.com/ess/kennel"
)

type Group struct{}

func (steps *Group) StepUp(s kennel.Suite) {
	s.Step(`^I have a group named mygroup$`, func() error {
		return nil
	})

	s.Step(`^I have a script that determines if I should scale up or down$`, func() error {
		return nil
	})

	s.Step(`^there is capacity for the group to upscale$`, func() error {
		return nil
	})

	s.Step(`^the group is scaled up$`, func() error {
		return nil
	})

	s.Step(`^there is not capacity for the group to upscale$`, func() error {
		return nil
	})

	s.Step(`^no changes are made$`, func() error {
		return nil
	})

}

func init() {
	kennel.Register(new(Group))
}
