package workflows

import (
	"fmt"

	"github.com/engineyard/eycore/core"
	"github.com/engineyard/eycore/users"

	"github.com/engineyard/scaley/group"
)

type ScalingAGroup struct {
	GroupName string
}

func (workflow *ScalingAGroup) perform() error {
	return asCurrentUser(workflow.scale)
}

func (workflow *ScalingAGroup) scale(api core.Client, current *users.Model) error {

	// get the current group definition
	currentGroup, err := group.ByName(workflow.GroupName)
	if err != nil {
		return fmt.Errorf("Could not resolve group: %s", workflow.GroupName)
	}

	// validate the group config
	err = group.Validate(currentGroup)
	if err != nil {
		return err
	}

	// block other attempts to scale the group
	err = group.Lock(currentGroup)
	if err != nil {
		return err
	}

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
		// If the ops are the same, Scale the group in the requested direction
		err := group.Scale(currentGroup, api, currentOp)
		if err != nil {
			return err
		}

		// Record the operation that we just performed
		group.RecordState(currentGroup, currentOp)
	} else {
		// If the ops are different, don't scale at all
		fmt.Println("Not scaling now.")
	}

	// unlock the group for alteration
	return group.Unlock(currentGroup)
}
