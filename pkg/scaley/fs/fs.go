package fs

import (
	"fmt"
	"os"

	"github.com/spf13/afero"
)

var Root = afero.NewOsFs()

func ReadFile(path string) ([]byte, error) {
	return afero.ReadFile(Root, path)
}

func FileExists(path string) bool {
	_, err := Root.Stat(path)

	if os.IsNotExist(err) {
		return false
	}

	return true
}

var CreateDir = func(path string) {
	if !FileExists(path) {
		err := Root.MkdirAll(path, 0755)
		if err != nil {
			fmt.Println("Could not create", path)
			os.Exit(255)
		}
	}
}
