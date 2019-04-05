package scaley

import (
	"github.com/ess/dry"
)

// ScalingEvent is a collection of data that acts as both input and output to
// a Scale transaction.
type ScalingEvent struct {
	GroupName   string
	Services    *Services
	Group       *Group
	Strategy    Strategy
	Direction   Direction
	Servers     []*Server
	Environment *Environment
	Candidates  []*Server
	Scaled      []*Server
	Failed      []*Server
	Error       error
}

func eventify(input dry.Value) *ScalingEvent {
	e := input.(*ScalingEvent)

	return e
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
