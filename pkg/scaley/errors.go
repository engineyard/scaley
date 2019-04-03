package scaley

import (
	"fmt"
)

type CannotLoadGroup struct {
	GroupName string
	Err       error
}

func (e CannotLoadGroup) Error() string {
	return fmt.Sprintf("%s could not be loaded: %s", e.GroupName, e.Err.Error())
}

type UnspecifiedScalingScript struct {
	Group *Group
}

func (e UnspecifiedScalingScript) Error() string {
	return fmt.Sprintf(
		"%s requires a scaling script",
		e.Group.Name,
	)
}

type MissingScalingScript struct {
	Group *Group
}

func (e MissingScalingScript) Error() string {
	return fmt.Sprintf(
		"%s scaling_script does not exist on the file system",
		e.Group.Name,
	)
}

type UnspecifiedScalingServers struct {
	Group *Group
}

func (e UnspecifiedScalingServers) Error() string {
	return fmt.Sprintf(
		"%s contains no scaling servers",
		e.Group.Name,
	)
}

type InvalidScalingServer struct {
	Group  *Group
	Server string
}

func (e InvalidScalingServer) Error() string {
	return fmt.Sprintf(
		"%s contains invlid scaling server %s",
		e.Group.Name,
		e.Server,
	)
}

type NoViableCandidates struct {
	Group     *Group
	Direction Direction
}

func (e NoViableCandidates) Error() string {
	return fmt.Sprintf(
		"%s has no viable candidates to scale %s",
		e.Group,
		e.Direction,
	)
}

type GroupIsLocked struct {
	Group *Group
}

func (e GroupIsLocked) Error() string {
	return fmt.Sprintf(
		"group operations are locked for %s",
		e.Group.Name,
	)
}

type LockFailure struct {
	Group *Group
}

func (e LockFailure) Error() string {
	return fmt.Sprintf(
		"could not lock group %s",
		e.Group.Name,
	)
}

type UnlockFailure struct {
	Group *Group
}

func (e UnlockFailure) Error() string {
	return fmt.Sprintf(
		"could not unlock group %s",
		e.Group.Name,
	)
}

type NoChangeRequired struct{}

func (e NoChangeRequired) Error() string {
	return "no scaling is necessary at this time"
}
