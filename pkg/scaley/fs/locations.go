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

// GroupConfigs returns the absolute path to the scaley group configuration
// directory after ensuring that the directory exists on Root
func GroupConfigs() string {
	path := fmt.Sprintf("%s/groups", configDir())
	CreateDir(path)
	return path
}

// Locks returns the absolute path to the scaley locks directory after ensuring
// that the directory exists on Root
func Locks() string {
	path := fmt.Sprintf("%s/lock", stateDir())
	CreateDir(path)
	return path
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

// CreateDir creates all elements of the given directory path on Root. If there
// are errors along the way, the program terminates via non-panic means.
var CreateDir = func(path string) {
	if !FileExists(path) {
		err := Root.MkdirAll(path, 0644)
		if err != nil {
			fmt.Println("Could not create", path)
			os.Exit(255)
		}
	}
}

// FileExists checks for the existence of the given path on Root. If the path
// exists, it returns true. Otherwise, it returns false.
func FileExists(path string) bool {
	_, err := Root.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	return true
}

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
