// Package fs provides services that interact directly with the local file
// system.
package fs

import (
	"github.com/spf13/afero"
)

// Root is an afero-wrapped filesystem that is considered the root fs for the
// application.
var Root = afero.NewOsFs()

// ReadFile attempts to read the contents of the given file path. If there are
// any errors along the way, an error is returned. Otherwise, the byte-
// representation of the file is returned.
func ReadFile(path string) ([]byte, error) {
	return afero.ReadFile(Root, path)
}
