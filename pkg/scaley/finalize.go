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

	// handle no-op
	if noOpDetected(err) {
		// do nothing, and return no error
		return nil
	}

	// handle scaling failure

	// pass all other errors upstream

	return err
}

func noOpDetected(err error) bool {
	_, c1 := err.(NoChangeRequired)

	return c1
}

func finalizeSuccess(result dry.Result) error {
	//event := eventify(result.Value())

	return errors.New("Unimplemented")
}
