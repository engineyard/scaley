package fs

import (
	"fmt"

	"github.com/engineyard/scaley/pkg/scaley"
)

// LockService is a service that knows how to lock and unlock groups via the
// file system.
type LockService struct{}

// NewLockService returns a new instance of LockService.
func NewLockService() *LockService {
	return &LockService{}
}

// Lock attempts to lock the given group. If there are issues in doing so, an
// error is returned. Otherwise, nil is returned.
func (service LockService) Lock(group *scaley.Group) error {
	if service.Locked(group) {
		return scaley.GroupIsLocked{Group: group}
	}

	l, err := Root.Create(lockfile(group))
	if err != nil {
		return scaley.LockFailure{Group: group}
	}

	l.Close()

	return nil
}

// Unlock attempts to unlock the given group. If there are issues in doing so,
// an error is returned. Otherwise, nil is returned.
func (service LockService) Unlock(group *scaley.Group) error {
	if service.Locked(group) {
		err := Root.Remove(lockfile(group))
		if err != nil {
			return scaley.UnlockFailure{Group: group}
		}
	}

	return nil
}

// Locked takes a group and returns a boolean that expresses whether or not that
// group is currently locked.
func (service LockService) Locked(group *scaley.Group) bool {
	return FileExists(lockfile(group))
}

func lockfile(group *scaley.Group) string {
	return fmt.Sprintf("%s/%s", Locks(), group.Name)
}

// Copyright © 2019 Engine Yard, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
