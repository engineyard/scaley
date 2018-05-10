package group

import (
	"github.com/engineyard/scaley/pkg/basher"
)

type scalingScript struct {
	group *Group
}

func ScalingScriptResult(group *Group) string {
	switch basher.Run(group.ScalingScript) {
	case 2:
		return "up"
	case 1:
		return "down"
	}

	return "noop"
}
