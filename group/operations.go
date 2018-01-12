package group

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/engineyard/eycore/core"
	"github.com/ess/spawning"
	"gopkg.in/yaml.v2"

	"github.com/engineyard/scaley/common"
)

func ByName(name string) (*Group, error) {
	var group *Group
	var err error

	dir := common.GroupConfigs()
	file := fmt.Sprintf("%s/%s.yml", dir, name)

	if !common.FileExists(file) {
		return nil, fmt.Errorf("No group named '%s'", name)
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(data, group)
	if err != nil {
		return nil, err
	}

	group.Name = name

	return group, nil
}

func Validate(group *Group) error {

	if len(group.ScalingScript) == 0 {
		return fmt.Errorf("The group requires a scaling_script")
	}

	if !common.FileExists(group.ScalingScript) {
		return fmt.Errorf("The group's scaling_script does not exist on the file system")
	}

	if len(group.PermanentServers) < 1 {
		return fmt.Errorf("The group contains no permanent servers")
	}

	if len(group.ScalingServers) < 1 {
		return fmt.Errorf("The group contains no scaling servers")
	}

	return nil
}

func Lock(group *Group) error {
	lockfile := lockfile(group)

	if common.FileExists(lockfile) {
		return fmt.Errorf("Group operations are locked for '%s'", group.Name)
	}

	l, err := os.Create(lockfile)
	if err != nil {
		return fmt.Errorf("Could not create lock file for '%s'", group.Name)
	}

	l.Close()

	return nil
}

func Unlock(group *Group) error {
	lockfile := lockfile(group)

	if common.FileExists(lockfile) {
		err := os.Remove(lockfile)
		if err != nil {
			return fmt.Errorf("Could not remove lock file for '%s'", group.Name)
		}
	}

	return nil
}

func lockfile(group *Group) string {
	return fmt.Sprintf("%s/%s", common.Locks(), group.Name)
}

func LastOperation(group *Group) string {
	lastfile := lastOpFile(group)

	if common.FileExists(lastfile) {
		data, err := ioutil.ReadFile(lastfile)
		if err != nil {
			output := fmt.Sprintf("Could not load last op for %s", group.Name)
			panic(output)
		}

		return string(data)
	}

	return "down"
}

func dataDir(group *Group) string {
	path := fmt.Sprintf("%s/%s", common.DataDir(), group.Name)
	common.CreateDir(path)
	return path
}
func lastOpFile(group *Group) string {
	return fmt.Sprintf("%s/lastop", dataDir(group))
}

func CurrentState(group *Group) string {
	statefile := stateFile(group)

	if common.FileExists(statefile) {
		data, err := ioutil.ReadFile(statefile)
		if err != nil {
			output := fmt.Sprintf("Could not load current state for %s", group.Name)
			panic(output)
		}

		return string(data)
	}

	return "down"
}

func stateFile(group *Group) string {
	return fmt.Sprintf("%s/state", dataDir(group))
}

func ScalingScriptResult(group *Group) string {
	result := spawning.Run(group.ScalingScript)

	if result.Success {
		return "up"
	}

	return "down"
}

func Scale(group *Group, api core.Client, direction string) error {
	panic("group.Scale() not implemented!")

	return nil
}

func RecordOp(group *Group, op string) error {
	lastfile := lastOpFile(group)

	err := ioutil.WriteFile(lastfile, []byte(op), 0644)
	if err != nil {
		return fmt.Errorf(
			"Could not write last op for %s: %s",
			group.Name,
			err.Error(),
		)
	}

	return nil
}

func RecordState(group *Group, state string) error {
	statefile := stateFile(group)

	err := ioutil.WriteFile(statefile, []byte(state), 0644)
	if err != nil {
		return fmt.Errorf(
			"Could not write state for %s: %s",
			group.Name,
			err.Error(),
		)
	}

	return nil
}
