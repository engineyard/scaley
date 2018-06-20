package workflows

import (
	"fmt"

	"github.com/engineyard/eycore/core"
	"github.com/engineyard/eycore/users"

	"github.com/engineyard/scaley/pkg/group"
	"github.com/engineyard/scaley/pkg/notifier"
	"github.com/engineyard/scaley/pkg/util"
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
	desiredOp := group.ScalingScriptResult(currentGroup)

	if desiredOp == "noop" {
		return nil
	}

	// provided that there's both popular demand and scaling capability
	if currentGroup.CanScale(desiredOp) {
		// Notify upstream that we're starting a scaling event
		notifier.Info(currentGroup, fmt.Sprintf("Scaling %s", desiredOp))

		// If the ops are the same, Scale the group in the requested direction
		err := group.Scale(currentGroup, api, desiredOp)
		if err != nil {
			// Notify upstream of the scaling failure
			notifier.Failure(
				currentGroup,
				fmt.Sprintf("Could not be scaled %s - %s", desiredOp, err.Error()),
			)

			return err
		}

		// Notify upstream of the scaling success
		notifier.Success(
			currentGroup,
			fmt.Sprintf("Successfully scaled %s", desiredOp),
		)

	} else {
		// The group can't be scaled in the desired direction
		if desiredOp == "up" {
			// If the desired outcome was an upscale event, log that we can't scale
			// and that the group might need to be bigger
			notifier.Info(
				currentGroup,
				"Cannot be scaled up - Consider adding more servers to the group",
			)
		}
	}

	return nil
}
