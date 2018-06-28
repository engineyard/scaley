package group

import (
	"math/rand"
	"testing"

	"github.com/ess/testscope"

	"github.com/engineyard/scaley/pkg/basher"
)

func TestScalingScriptResult(t *testing.T) {
	testscope.SkipUnlessUnit(t)

	g := &Group{ScalingScript: "dummy"}

	status := 0
	basher.Run = func(command string) int {
		return status
	}

	t.Run("when the scaling script exits with status 1", func(t *testing.T) {
		status = 1
		result := ScalingScriptResult(g)

		t.Run("the desired op is 'down'", func(t *testing.T) {
			if result != "down" {
				t.Errorf("expected a downscale event, got %s", result)
			}
		})
	})

	t.Run("when the scaling script exits with status 2", func(t *testing.T) {
		status = 2
		result := ScalingScriptResult(g)

		t.Run("the desired op is 'up'", func(t *testing.T) {
			if result != "up" {
				t.Errorf("expected an upscale event, got %s", result)
			}
		})
	})

	t.Run("when the scaling script exits with status 0", func(t *testing.T) {
		status = 0
		result := ScalingScriptResult(g)

		t.Run("the desired op is 'noop'", func(t *testing.T) {
			if result != "noop" {
				t.Errorf("expected a noop, got %s", result)
			}
		})
	})

	t.Run("when the scaling script exits with unknown status", func(t *testing.T) {
		status = rand.Int() + 2
		result := ScalingScriptResult(g)

		t.Run("the desired op is 'noop'", func(t *testing.T) {
			if result != "noop" {
				t.Errorf("expected a noop, got %s", result)
			}
		})
	})
}
