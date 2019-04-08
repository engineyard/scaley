package scaley

// Direction describes the desired direction (up, down, none) for a scaling
// event.
type Direction int

// String returns a text representation of a direction.
func (d Direction) String() string {
	switch d {
	case Up:
		return "UP"
	case Down:
		return "DOWN"
	default:
		return "NONE"
	}
}

// DesiredState returns the state that a server should be in for scaling to
// be considered for a given direction.
func (d Direction) DesiredState() string {
	if d == Up {
		return "stopped"
	}

	return "running"
}

const (
	// None indicates that no scaling event should occur.
	None Direction = iota
	// Down indicates that the group should be scaled down.
	Down
	// Up indicates that the group should be scaled up.
	Up
)

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
