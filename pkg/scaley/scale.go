package scaley

import (
	"fmt"

	"github.com/ess/debuggable"
	"github.com/ess/dry"
)

// Scale is a factory that produces a new Transaction that can be called
// with Call.
func Scale() *dry.Transaction {
	return dry.NewTransaction(
		loadGroup,
		validateGroup,
		lockGroup,
		loadServers,
		loadEnvironment,
		calculateDirection,
		calculateCandidates,
		announceScalingStart,
		scaleCandidates,
		configureEnvironment,
	)
}

func loadGroup(input dry.Value) dry.Result {
	event := eventify(input)

	name := event.GroupName

	group, err := event.Services.Groups.Get(name)
	if err != nil {
		event.Error = CannotLoadGroup{name, err}
		return dry.Failure(event)
	}

	event.Group = group
	event.Strategy = CalculateStrategy(group)

	return dry.Success(event)
}

func validateGroup(input dry.Value) dry.Result {
	event := eventify(input)

	group := event.Group

	// verify that the group specifies a scaling script
	if len(group.ScalingScript) == 0 {
		event.Error = UnspecifiedScalingScript{group}

		return dry.Failure(event)
	}

	// verify that the scaling script exists

	if !event.Services.Scripts.Exists(group.ScalingScript) {
		event.Error = MissingScalingScript{group}

		return dry.Failure(event)
	}

	// verify that the group contains scaling servers

	if len(group.ScalingServers) < 1 {
		event.Error = UnspecifiedScalingServers{group}

		return dry.Failure(event)
	}

	// All checks passed, continue the transaction
	return dry.Success(event)
}

func lockGroup(input dry.Value) dry.Result {
	event := eventify(input)
	group := event.Group
	locker := event.Services.Locker

	if locker.Locked(group) {
		event.Error = GroupIsLocked{group}

		return dry.Failure(event)
	}

	if err := event.Services.Locker.Lock(group); err != nil {
		event.Error = LockFailure{group}

		return dry.Failure(event)
	}

	return dry.Success(event)
}

func loadServers(input dry.Value) dry.Result {
	event := eventify(input)
	group := event.Group

	event.Servers = make([]*Server, 0)

	for _, id := range group.ScalingServers {
		server, err := event.Services.Servers.Get(id)
		if err != nil {
			event.Error = InvalidScalingServer{group, id}

			return dry.Failure(event)
		}

		event.Servers = append(event.Servers, server)
	}

	return dry.Success(event)
}

func loadEnvironment(input dry.Value) dry.Result {
	event := eventify(input)
	server := event.Servers[0]

	environment, err := event.Services.Environments.Get(server.EnvironmentID)
	if err != nil {
		event.Error = InvalidEnvironment{event.Group}

		return dry.Failure(event)
	}

	event.Environment = environment

	return dry.Success(event)
}

func calculateDirection(input dry.Value) dry.Result {
	event := eventify(input)
	script := event.Group.ScalingScript

	event.Direction = Direction(event.Services.Runner.Run(script))

	if event.Direction == None {
		event.Error = NoChangeRequired{}

		return dry.Failure(event)
	}

	return dry.Success(event)
}

func announceScalingStart(input dry.Value) dry.Result {
	event := eventify(input)

	event.Services.Log.Info(
		event.Group,
		fmt.Sprintf("Scaling %s", event.Direction),
	)

	return dry.Success(event)
}

func calculateCandidates(input dry.Value) dry.Result {
	event := eventify(input)

	for _, server := range event.Servers {
		if server.State == event.Direction.DesiredState() {
			event.Candidates = append(event.Candidates, server)
		}
	}

	if len(event.Candidates) < 1 {
		event.Error = NoViableCandidates{
			event.Group,
			event.Direction,
		}

		return dry.Failure(event)
	}

	return dry.Success(event)
}

func scaleCandidates(input dry.Value) dry.Result {
	event := eventify(input)
	toScale := make([]*Server, 0)

	var method func(*Server, *ScalingEvent) error

	switch event.Strategy {
	case Individual:
		toScale = append(toScale, event.Candidates[0])
	default:
		toScale = event.Candidates
	}

	switch event.Direction {
	case Up:
		method = scaleCandidateUp
	default:
		method = scaleCandidateDown
	}

	for _, server := range toScale {
		err := method(server, event)
		if err != nil {
			if debuggable.Enabled() {
				fmt.Println("[scaley debug] server state change error:", err)
			}
			event.Failed = append(event.Failed, server)
		} else {
			event.Scaled = append(event.Scaled, server)
		}
	}

	if len(event.Failed) > 0 {
		event.Error = ScalingFailure{
			event.Group,
			event.Direction,
			event.Scaled,
			event.Failed,
		}

		return dry.Failure(event)
	}

	return dry.Success(event)
}

func scaleCandidateUp(candidate *Server, event *ScalingEvent) error {
	return event.Services.Servers.Start(candidate)
}

func scaleCandidateDown(candidate *Server, event *ScalingEvent) error {
	stopscript := event.Group.StopScript

	if len(stopscript) > 0 {
		command := fmt.Sprintf("%s %s", stopscript, candidate.Hostname)

		ssr := event.Services.Runner.Run(command)

		if !event.Group.IgnoreStopScriptErrors {
			if ssr != 0 {
				return fmt.Errorf("stop script failure")
			}
		}
	}

	return event.Services.Servers.Stop(candidate)
}

func configureEnvironment(input dry.Value) dry.Result {
	event := eventify(input)

	err := event.Services.Environments.Configure(event.Environment)
	if err != nil {
		event.Error = ChefFailure{
			event.Group,
			event.Environment,
		}

		return dry.Failure(event)
	}

	return dry.Success(event)
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
