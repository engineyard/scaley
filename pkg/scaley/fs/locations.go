package fs

import (
	"fmt"
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
