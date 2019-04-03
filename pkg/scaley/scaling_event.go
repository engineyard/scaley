package scaley

import (
	"github.com/ess/dry"
)

type ScalingEvent struct {
	GroupName  string
	Services   *Services
	Group      *Group
	Strategy   Strategy
	Direction  Direction
	Servers    []*Server
	Candidates []*Server
	Scaled     []*Server
	Failed     []*Server
	Error      error
}

func eventify(input dry.Value) *ScalingEvent {
	e := input.(*ScalingEvent)

	return e
}
