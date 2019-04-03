package scaley

import (
	"github.com/ess/dry"
)

// Scale is a factory that produces a new Transaction that can be called
// with Call.
func Scale() *dry.Transaction {
	return dry.NewTransaction(
		loadGroup,
		validateGroup,
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

	return dry.Failure(event)
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

	// verify that the group contains scaling servers

	if len(group.ScalingServers) < 1 {
		event.Error = UnspecifiedScalingServers{group}

		return dry.Failure(event)
	}

	// All checks passed, continue the transaction
	return dry.Success(event)
}

func loadServers(input dry.Value) dry.Result {
	event := eventify(input)

	return dry.Failure(event)
}

func calculateDirection(input dry.Value) dry.Result {
	event := eventify(input)

	// execute the scaling script, set the direction based on it.
	// If the direction is no-op, fail

	return dry.Failure(event)
}

func announceScalingStart(input dry.Value) dry.Result {
	event := eventify(input)

	return dry.Failure(event)
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

	return dry.Failure(event)
}

func configureEnvironment(input dry.Value) dry.Result {
	event := eventify(input)

	return dry.Failure(event)
}
