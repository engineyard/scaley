package fs

import (
	"github.com/spf13/afero"
)

var Root = afero.NewOsFs()

func ReadFile(path string) ([]byte, error) {
	return afero.ReadFile(Root, path)
}
