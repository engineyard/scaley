package scaley

import (
	"github.com/ess/dry"
)

// ScalingEvent is a collection of data that acts as both input and output to
// a Scale transaction.
type ScalingEvent struct {
	GroupName   string
	Services    *Services
	Group       *Group
	Strategy    Strategy
	Direction   Direction
	Servers     []*Server
	Environment *Environment
	Candidates  []*Server
	Scaled      []*Server
	Failed      []*Server
	Error       error
}

func eventify(input dry.Value) *ScalingEvent {
	e := input.(*ScalingEvent)

	return e
}
