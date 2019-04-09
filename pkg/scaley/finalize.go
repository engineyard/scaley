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
	group := event.Group

	// handle no-op
	if noOpDetected(err) {
		// unlock the group
		lerr := locker.Unlock(group)
		if lerr != nil {
			logUnlockFailure(log, group)
		}

		// return no error
		return nil
	}

	direction := strings.ToLower(event.Direction.String())

	// handle insufficient capacity
	if insufficientCapacityDetected(err) {
		if event.Direction == Up {
			log.Info(
				group,
				"Cannot be scaled up - Consider adding more servers to the group",
			)
		}

		lerr := locker.Unlock(group)
		if lerr != nil {
			logUnlockFailure(log, group)
		}

		return nil
	}

	// handle scaling )ailure
	if scalingFailureDetected(err) {
		action := "starting"

		if event.Direction == Down {
			action = "stopping"
		}

		log.Failure(
			group,
			fmt.Sprintf(
				"Could not be scaled %s - Errors occurred while %s servers, please contact support. %s",
				direction,
				action,
				failureDetails(event.Scaled, event.Failed),
			),
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
			fmt.Sprintf(
				"Could not be scaled %s - A Chef error occurred while %sscaling the group. Please contact support. %s",
				direction,
				direction,
				failureDetails(event.Scaled, event.Failed),
			),
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
	scaled := make([]string, 0)

	for _, server := range event.Scaled {
		scaled = append(scaled, server.ProvisionedID)
	}

	log.Success(
		group,
		fmt.Sprintf(
			"Successfully scaled %s. Servers affected: %s",
			event.Direction.String(),
			strings.Join(scaled, ", "),
		),
	)

	return nil
}

func logUnlockFailure(log LogService, group *Group) {
	log.Failure(group, "Cannot unlock the group. Please contact support.")
}

func failureDetails(scaled []*Server, failed []*Server) string {
	failedIDs := make([]string, 0)
	for _, f := range failed {
		failedIDs = append(failedIDs, f.ProvisionedID)
	}

	scaledIDs := make([]string, 0)
	for _, s := range scaled {
		scaledIDs = append(scaledIDs, s.ProvisionedID)
	}

	parts := make([]string, 0)

	if len(failedIDs) > 0 {
		parts = append(
			parts,
			"(Failed to start/stop: "+strings.Join(failedIDs, ", ")+")",
		)
	}

	if len(scaledIDs) > 0 {
		parts = append(
			parts,
			"(Successfully started/stopped: "+strings.Join(scaledIDs, ", ")+")",
		)
	}

	return strings.Join(parts, " ")
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
