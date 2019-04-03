package scaley

import (
	"strings"
)

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
	Legion Strategy = iota
	Individual
)

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
