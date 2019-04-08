package fs

import (
	"fmt"
	"os"

	"github.com/spf13/afero"
)

// Root is the abstracted file system used for all operations in this package.
var Root = afero.NewOsFs()

// ReadFile takes a path, reads the file from the file system, and returns its
// contents as a byte array as well as an error. If there are issues along the
// way, the error is populated. Otherwise, the error is nil.
func ReadFile(path string) ([]byte, error) {
	return afero.ReadFile(Root, path)
}

// FileExists takes a path, checks to see if that file exists on the file
// system. If the file does exist, it returns true. Otherwise, it returns false.
func FileExists(path string) bool {
	_, err := Root.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	return true
}

// CreateDir takes a path and attempts to create it and all of its ancestors
// on the file system. If there are problems along the way, the program exits
// with code 255.
var CreateDir = func(path string) {
	if !FileExists(path) {
		err := Root.MkdirAll(path, 0755)
		if err != nil {
			fmt.Println("Could not create", path)
			os.Exit(255)
		}
	}
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
