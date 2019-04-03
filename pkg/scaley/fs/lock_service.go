package fs

import (
	"fmt"

	"github.com/engineyard/scaley/pkg/scaley"
)

type LockService struct{}

func NewLockService() *LockService {
	return &LockService{}
}

func (service LockService) Lock(group *scaley.Group) error {
	if service.Locked(group) {
		return scaley.GroupIsLocked{group}
	}

	l, err := Root.Create(lockfile(group))
	if err != nil {
		return scaley.LockFailure{group}
	}

	l.Close()

	return nil
}

func (service LockService) Unlock(group *scaley.Group) error {
	if service.Locked(group) {
		err := Root.Remove(lockfile(group))
		if err != nil {
			return scaley.UnlockFailure{group}
		}
	}

	return nil
}

func (service LockService) Locked(group *scaley.Group) bool {
	return FileExists(lockfile(group))
}

func lockfile(group *scaley.Group) string {
	return fmt.Sprintf("%s/%s", Locks(), group.Name)
}
