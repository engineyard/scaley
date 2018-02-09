package group

import (
	"fmt"
	"io/ioutil"

	"github.com/engineyard/scaley/common"
)

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

func stateFile(group *Group) string {
	return fmt.Sprintf("%s/state", dataDir(group))
}
