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

package fs

import (
	"fmt"
	"os"
)

const (
	config = "/etc/scaley"
	data   = "/var/lib/scaley"
	state  = "/var/run/scaley"
)

func GroupConfigs() string {
	path := fmt.Sprintf("%s/groups", configDir())
	CreateDir(path)
	return path
}

func Locks() string {
	path := fmt.Sprintf("%s/lock", stateDir())
	CreateDir(path)
	return path
}

func DataDir() string {
	return dataDir()
}

func stateDir() string {
	path := state
	CreateDir(path)
	return path
}

func configDir() string {
	path := config
	CreateDir(path)
	return path
}

func dataDir() string {
	path := data
	CreateDir(path)
	return path
}

var CreateDir = func(path string) {
	if !FileExists(path) {
		err := Root.MkdirAll(path, 0644)
		if err != nil {
			fmt.Println("Could not create", path)
			os.Exit(255)
		}
	}
}

func FileExists(path string) bool {
	_, err := Root.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	return true
}
