package scaley

import (
	"fmt"
	"strings"

	"github.com/ess/dry"
)

// Finalize is the last step of all scaley runs. It takes a dry.Result with an
// embedded scaling event and handles things like cleanup, logging of
// event completion, and so on. If the scaling event failed, an error is
// returned. Otherwise, nil is returned.
func Finalize(result dry.Result) error {
	if result.Failure() {
		return finalizeFailure(result)
	}

	return finalizeSuccess(result)
}

func finalizeFailure(result dry.Result) error {
	event := eventify(result.Error())
	log := event.Services.Log
	locker := event.Services.Locker
	err := event.Error

	// handle no-op
	if noOpDetected(err) {
		// do nothing, and return no error
		return nil
	}

	group := event.Group
	direction := strings.ToLower(event.Direction.String())

	// handle insufficient capacity
	if insufficientCapacityDetected(err) {
		if event.Direction == Up {
			log.Info(
				group,
				"Cannot be scaled up - Consider adding more servers to the group",
			)
		}

		return nil
	}

	// handle scaling )ailure
	if scalingFailureDetected(err) {
		action := "starting"

		if event.Direction == Down {
			action = "stopping"
		}

		failures := make([]string, 0)
		for _, f := range event.Failed {
			failures = append(failures, f.ProvisionedID)
		}

		log.Failure(
			group,
			fmt.Sprintf("Could not be scaled %s - Errors occurred while %s these servers, please contact support: %s", direction, action, strings.Join(failures, ", ")),
		)

		lerr := locker.Unlock(group)
		if lerr != nil {
			logUnlockFailure(log, group)
		}

		return err
	}

	// handle chef failure
	if chefFailureDetected(err) {
		log.Failure(
			group,
			fmt.Sprintf("Could not be scaled %s - A Chef error occurred while %sscaling the group. Please contact support.", direction, direction),
		)
	}

	// pass all other errors upstream

	return err
}

func noOpDetected(err error) bool {
	_, c1 := err.(NoChangeRequired)

	return c1
}

func scalingFailureDetected(err error) bool {
	_, c1 := err.(ScalingFailure)

	return c1
}

func chefFailureDetected(err error) bool {
	_, c1 := err.(ChefFailure)

	return c1
}

func insufficientCapacityDetected(err error) bool {
	_, c1 := err.(NoViableCandidates)

	return c1
}

func finalizeSuccess(result dry.Result) error {
	event := eventify(result.Value())
	log := event.Services.Log
	locker := event.Services.Locker
	group := event.Group

	// unlock the group
	err := locker.Unlock(group)
	if err != nil {
		logUnlockFailure(log, group)

		return err
	}

	// log success
	log.Success(
		group,
		fmt.Sprintf(
			"Successfully scaled %s",
			event.Direction.String(),
		),
	)

	return nil
}

func logUnlockFailure(log LogService, group *Group) {
	log.Failure(group, "Cannot unlock the group. Please contact support.")
}
