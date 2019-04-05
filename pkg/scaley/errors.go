package scaley

import (
	"fmt"
)

// CannotLoadGroup is an error that is raised when the group data cannot be
// read from its data source.
type CannotLoadGroup struct {
	GroupName string
	Err       error
}

func (e CannotLoadGroup) Error() string {
	return fmt.Sprintf("%s could not be loaded: %s", e.GroupName, e.Err.Error())
}

// UnspecifiedScalingScript is an error that is raised when a group does not
// have a scaling script configured.
type UnspecifiedScalingScript struct {
	Group *Group
}

func (e UnspecifiedScalingScript) Error() string {
	return fmt.Sprintf(
		"%s requires a scaling_script",
		e.Group.Name,
	)
}

// MissingScalingScript is an error that is raised when a group's configured
// scaling script does not actually exist.
type MissingScalingScript struct {
	Group *Group
}

func (e MissingScalingScript) Error() string {
	return fmt.Sprintf(
		"%s scaling_script does not exist on the file system",
		e.Group.Name,
	)
}

// UnspecifiedScalingServers is an error that is raised when a group has no
// scaling servers configured.
type UnspecifiedScalingServers struct {
	Group *Group
}

func (e UnspecifiedScalingServers) Error() string {
	return fmt.Sprintf(
		"%s contains no scaling servers",
		e.Group.Name,
	)
}

// InvalidScalingServer is an error that is raised when a group contains a
// scaling server that cannot be found on the upstream Engine Yard API.
type InvalidScalingServer struct {
	Group  *Group
	Server string
}

func (e InvalidScalingServer) Error() string {
	return fmt.Sprintf(
		"%s contains invalid scaling server %s",
		e.Group.Name,
		e.Server,
	)
}

// NoViableCandidates is an error that is raised when there are no servers in
// the desired state for the associated scaling event.
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

// GroupIsLocked is an error that is raised when the group associated with a
// scaling event is already locked for scaling operations.
type GroupIsLocked struct {
	Group *Group
}

func (e GroupIsLocked) Error() string {
	return fmt.Sprintf(
		"group operations are locked for %s",
		e.Group.Name,
	)
}

// LockFailure is an error that is raised when an attempt to lock a group fails.
type LockFailure struct {
	Group *Group
}

func (e LockFailure) Error() string {
	return fmt.Sprintf(
		"could not lock group %s",
		e.Group.Name,
	)
}

// UnlockFailure is an error that is raised when an attempt to unlock a group
// fails.
type UnlockFailure struct {
	Group *Group
}

func (e UnlockFailure) Error() string {
	return fmt.Sprintf(
		"could not unlock group %s",
		e.Group.Name,
	)
}

// NoChangeRequired is an error that is raised when the scaling script for a
// scaling script's group indicates that the group should not be scaled.
type NoChangeRequired struct{}

func (e NoChangeRequired) Error() string {
	return "no scaling is necessary at this time"
}

// ScalingFailure is an error that is raised when an attempt to scale a group's
// scaling servers up or down fails.
type ScalingFailure struct {
	Group     *Group
	Direction Direction
	Scaled    []*Server
	Failed    []*Server
}

func (e ScalingFailure) Error() string {
	return fmt.Sprintf(
		"could not be scaled %s",
		e.Direction.String(),
	)
}

// InvalidEnvironment is an error that is raised when the environment associated
// with a scaling script's group cannot be loaded from the Engine Yard API.
type InvalidEnvironment struct {
	Group *Group
}

func (e InvalidEnvironment) Error() string {
	return fmt.Sprintf(
		"could not load environment for %s",
		e.Group.Name,
	)
}

// ChefFailure is an error that is raised when an attempt to reconfigure an
// environment fails.
type ChefFailure struct {
	Group       *Group
	Environment *Environment
}

func (e ChefFailure) Error() string {
	return fmt.Sprintf(
		"could not configure %s for group %s",
		e.Environment.Name,
		e.Group.Name,
	)
}

// Copyright © 2019 Engine Yard, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
