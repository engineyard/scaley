package notifier

import (
	"fmt"

	"github.com/engineyard/scaley/pkg/group"
)

func Failure(g *group.Group, message string) {
	notify(
		failure,
		normalize(g, message),
	)
}

func Info(g *group.Group, message string) {
	notify(
		warning,
		normalize(g, message),
	)
}

func Success(g *group.Group, message string) {
	notify(
		okay,
		normalize(g, message),
	)
}

func normalize(g *group.Group, message string) string {
	return fmt.Sprintf("Group[%s]: %s", g.Name, message)
}
