package group

import (
	"fmt"

	"github.com/engineyard/scaley/pkg/common"
)

func Validate(group *Group) error {

	if len(group.ScalingScript) == 0 {
		return fmt.Errorf("%s requires a scaling_script", group.Name)
	}

	if !common.FileExists(group.ScalingScript) {
		return fmt.Errorf("%s scaling_script does not exist on the file system", group.Name)
	}

	if len(group.ScalingServers) < 1 {
		return fmt.Errorf("%s contains no scaling servers", group.Name)
	}

	for _, server := range group.ScalingServers {
		if server.Instance == nil {
			return fmt.Errorf("%s contains invalid scaling server %s", group.Name, server.ID)
		}
	}

	return nil
}
