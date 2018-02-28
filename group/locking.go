package group

import (
	"fmt"
	"os"

	"github.com/engineyard/eycore/core"

	"github.com/engineyard/scaley/common"
)

type locked func(group *Group, api core.Client) error

func LockedProcedure(group *Group, api core.Client, process locked) error {
	err := lock(group)
	if err != nil {
		return err
	}

	err = process(group, api)

	unlockErr := unlock(group)

	if err != nil {
		return err
	}

	return unlockErr
}

var lock = func(group *Group) error {
	lockfile := lockfile(group)

	if common.FileExists(lockfile) {
		return fmt.Errorf("Group operations are locked for '%s'", group.Name)
	}

	l, err := os.Create(lockfile)
	if err != nil {
		return fmt.Errorf("Could not create lock file for '%s'", group.Name)
	}

	l.Close()

	return nil
}

var unlock = func(group *Group) error {
	lockfile := lockfile(group)

	if common.FileExists(lockfile) {
		err := os.Remove(lockfile)
		if err != nil {
			return fmt.Errorf("Could not remove lock file for '%s'", group.Name)
		}
	}

	return nil
}

func lockfile(group *Group) string {
	return fmt.Sprintf("%s/%s", common.Locks(), group.Name)
}
