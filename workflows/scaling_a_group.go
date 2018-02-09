package workflows

import (
	"fmt"

	"github.com/engineyard/eycore/core"
	"github.com/engineyard/eycore/users"

	"github.com/engineyard/scaley/group"
	"github.com/engineyard/scaley/notifier"
	"github.com/engineyard/scaley/util"
)

type ScalingAGroup struct {
	GroupName string
}

func (workflow *ScalingAGroup) perform() error {
	return asCurrentUser(workflow.scale)
}

func (workflow *ScalingAGroup) scale(api core.Client, current *users.Model) error {

	// Ensure that there is a reporting URL in the config, as it is needed
	// for logging scaling events to upstream.
	err := util.CheckReportingURL()
	if err != nil {
		return err
	}

	// get the current group definition
	currentGroup, err := group.ByName(api, workflow.GroupName)
	if err != nil {
		return fmt.Errorf("Could not resolve group: %s", workflow.GroupName)
	}

	// validate the group config
	err = group.Validate(currentGroup)
	if err != nil {
		return err
	}

	return group.LockedProcedure(currentGroup, api, workflow.withLocking)
}

func (workflow *ScalingAGroup) withLocking(currentGroup *group.Group, api core.Client) error {
	// recall the last recorded scaling operation
	lastOp := group.LastOperation(currentGroup)

	// calculate the current operation
	currentOp := group.ScalingScriptResult(currentGroup)

	// get the current state of the group
	currentState := group.CurrentState(currentGroup)

	// record the current operation
	group.RecordOp(currentGroup, currentOp)

	// determine the opersation to actually perform
	if lastOp == currentOp && currentOp != currentState {
		// Notify upstream that we're starting a scaling event
		notifier.Info(currentGroup, fmt.Sprintf("Scaling %s", currentOp))

		// If the ops are the same, Scale the group in the requested direction
		err := group.Scale(currentGroup, api, currentOp)
		if err != nil {
			// Notify upstream of the scaling failure
			notifier.Failure(
				currentGroup,
				fmt.Sprintf("Could not be scaled %s - %s", currentOp, err.Error()),
			)

			return err
		}

		// Notify upstream of the scaling success
		notifier.Success(
			currentGroup,
			fmt.Sprintf("Successfully scaled %s", currentOp),
		)

		// Record the operation that we just performed
		group.RecordState(currentGroup, currentOp)
	} else {
		// If the ops are different, don't scale at all
		fmt.Println("Not scaling now.")
	}

	return nil
}
