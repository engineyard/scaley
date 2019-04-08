package fs

import (
	"fmt"
)

const (
	config = "/etc/scaley"
	data   = "/var/lib/scaley"
	state  = "/var/run/scaley"
)

// GroupConfigs is the file system location in which group definition files are
// stored.
func GroupConfigs() string {
	path := fmt.Sprintf("%s/groups", configDir())
	CreateDir(path)
	return path
}

// Locks is the file system location in which group lock files are stored.
func Locks() string {
	path := fmt.Sprintf("%s/lock", stateDir())
	CreateDir(path)
	return path
}

// DataDir is the file system location in which runtime data is stored.
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
