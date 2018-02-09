package group

import (
	"github.com/ess/spawning"
)

func ScalingScriptResult(group *Group) string {
	result := spawning.Run(group.ScalingScript)

	if result.Success {
		return "up"
	}

	return "down"
}
