package scaley

import (
	"strings"
)

// Strategy describes one of several possible scaling strategies for a group.
type Strategy int

func (s Strategy) String() string {
	switch s {
	case Individual:
		return "Individual"
	default:
		return "Legion"
	}
}

const (
	// Legion is an all-or-nothing scaling strategy that acts upon all scaling
	// servers in a group.
	Legion Strategy = iota
	// Individual is a conservative strategy that acts upon only a single scaling
	// server in a group.
	Individual
)

// CalculateStrategy takes a group and returns the strategy with which it is
// configured. If the group lacks a strategy, Legion is returned by default.
func CalculateStrategy(group *Group) Strategy {
	name := normalizedStrategyName(group.Strategy)

	switch name {
	case "individual":
		return Individual
	default:
		return Legion
	}
}

func normalizedStrategyName(name string) string {
	return strings.ToLower(
		strings.TrimSpace(
			name,
		),
	)
}
