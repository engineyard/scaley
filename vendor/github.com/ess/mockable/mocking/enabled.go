package mocking

import (
	"testing"

	"github.com/ess/mockable/core"
)

func EnabledDo(t *testing.T, procedure func(*testing.T)) {
	originallyMocked := core.Mocked()
	if !originallyMocked {
		core.Enable()
	}

	defer func() {
		if !originallyMocked {
			core.Disable()
		}

		if r := recover(); r != nil {
			t.Error("Mocked procedure panic:", r)
		}

	}()

	procedure(t)
}
