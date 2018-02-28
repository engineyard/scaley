package group

import (
	"fmt"
	"io/ioutil"

	"github.com/engineyard/scaley/common"
)

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

func dataDir(group *Group) string {
	path := fmt.Sprintf("%s/%s", common.DataDir(), group.Name)
	common.CreateDir(path)
	return path
}
func lastOpFile(group *Group) string {
	return fmt.Sprintf("%s/lastop", dataDir(group))
}
