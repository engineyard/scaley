// Package scaley describes the high-level domain for scaling groups of
// servers on the Engine Yard Cloud platform.
package scaley

type Direction int

const (
	None Direction = iota
	Down
	Up
)

type ServerState int

const (
	Unknown ServerState = iota
	Stopped
	Running
)

type Severity int

const (
	Okay Severity = iota
	Warning
	Failure
)

func (s Severity) String() {
	var severity string

	switch s {
	case Warning:
		severity = "WARNING"
	case Failure:
		severity = "FAILURE"
	case Okay:
		severity = "OKAY"
	default:
		severity = "UNKNOWN"
	}

	return severity
}

// Group is a representation of an autoscaling group
type Group struct {
	Name          string
	Servers       []Server
	ScalingScript string
	StopScript    string
	Strategy      string
	Environment   Environment
}

// Server is a representation of a server within a Group
type Server struct {
	ID            int
	ProvisionedID string
	State         ServerState
	EnvironmentID string
}

// Environment is a representation of an environment on the Engine Yard Platform
type Environment struct {
	ID   string
	Name string
}

// Strategy provides the high-level functionality for scaling a group
type Strategy interface {
	Upscale() error
	Downscale() error
}

// GroupService provides retrieval functionality for groups
type GroupService interface {
	Get(string) (Group, error)
}

// ServerService provides retrieval functionality for servers
type ServerService interface {
	Get(string) (Server, error)
}

// EnvironmentService provides retrieval functionality for environments
type EnvironmentService interface {
	Get(string) (Environment, error)
}

// OpsService provides low-level functionality related to scaling groups.
//
// More than anything, an OpsService's focus is around starting/stopping
// servers, configuring environments, and so on.
type OpsService interface {
	Start(Server) error
	Stop(Server) error
	Configure(Environment) error
}

type CandidateService interface {
	Single(Direction) Server
	All(Direction) []Server
}

type Logger interface {
	Info(Group, string)
	Success(Group, string)
	Failure(Group, string)
}

// Copyright © 2018 Engine Yard, Inc.
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
