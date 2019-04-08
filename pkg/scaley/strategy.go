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

// Copyright © 2019 Engine Yard, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
