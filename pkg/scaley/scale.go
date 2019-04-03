package scaley

import (
	"fmt"

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
		calculateDirection,
		announceScalingStart,
		calculateCandidates,
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

	var method func(*Server) error

	switch event.Strategy {
	case Individual:
		toScale = append(toScale, event.Candidates[0])
	default:
		toScale = event.Candidates
	}

	switch event.Direction {
	case Up:
		method = event.Services.Servers.Start
	default:
		method = event.Services.Servers.Stop
	}

	for _, server := range toScale {
		err := method(server)
		if err != nil {
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

func configureEnvironment(input dry.Value) dry.Result {
	event := eventify(input)

	return dry.Failure(event)
}
