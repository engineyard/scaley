package scaley

import (
	"errors"

	"github.com/ess/dry"
)

func Finalize(result dry.Result) error {
	if result.Failure() {
		return finalizeFailure(result)
	}

	return finalizeSuccess(result)
}

func finalizeFailure(result dry.Result) error {
	event := eventify(result.Error())
	err := event.Error

	// handle invalid group
	if invalidGroupDetected(err) {
		// return the error without logging
		return err
	}

	// handle no-op
	if noOpDetected(err) {
		// do nothing, and return no error
		return nil
	}

	// handle scaling failure

	return errors.New("Unimplemented")
}

func invalidGroupDetected(err error) bool {
	_, c1 := err.(UnspecifiedScalingScript)
	_, c2 := err.(MissingScalingScript)
	_, c3 := err.(UnspecifiedScalingServers)

	return c1 || c2 || c3
}

func noOpDetected(err error) bool {
	_, c1 := err.(NoChangeRequired)

	return c1
}

func finalizeSuccess(result dry.Result) error {
	//event := eventify(result.Value())

	return errors.New("Unimplemented")
}
