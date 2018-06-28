// Copyright © 2018 Engine Yard, Inc.
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

package group

import (
	"fmt"

	"github.com/engineyard/eycore/core"

	"github.com/engineyard/scaley/pkg/common"
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

	l, err := common.Root.Create(lockfile)
	if err != nil {
		return fmt.Errorf("Could not create lock file for '%s'", group.Name)
	}

	l.Close()

	return nil
}

var unlock = func(group *Group) error {
	lockfile := lockfile(group)

	if common.FileExists(lockfile) {
		err := common.Root.Remove(lockfile)
		if err != nil {
			return fmt.Errorf("Could not remove lock file for '%s'", group.Name)
		}
	}

	return nil
}

func lockfile(group *Group) string {
	return fmt.Sprintf("%s/%s", common.Locks(), group.Name)
}
