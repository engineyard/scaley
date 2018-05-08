package mocking

import (
	"testing"

	"github.com/ess/mockable/core"
)

func DisabledDo(t *testing.T, procedure func(*testing.T)) {
	originallyMocked := core.Mocked()
	if originallyMocked {
		core.Disable()
	}

	defer func() {
		if originallyMocked {
			core.Enable()
		}

		if r := recover(); r != nil {
			t.Error("Mocked procedure panic:", r)
		}

	}()

	procedure(t)

}
