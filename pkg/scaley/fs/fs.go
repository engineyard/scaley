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
