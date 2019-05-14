// Package scaley provides the domain objects, interfaces, and workflows for
// scaling server groups on Engine Yard Cloud.
package scaley

// Group is a representation of a scaling group.
type Group struct {
	Name                   string   `yaml:"name"`
	ScalingServers         []string `yaml:"scaling_servers"`
	ScalingScript          string   `yaml:"scaling_script"`
	StopScript             string   `yaml:"stop_script"`
	IgnoreStopScriptErrors bool     `yaml:"ignore_stop_script_errors"`
	UnlockOnFailure        bool     `yaml:"unlock_on_failure"`
	Strategy               string   `yaml:"strategy"`
}

// Server is a representation of a server in an Engine Yard Cloud environment.
type Server struct {
	ID            int
	ProvisionedID string
	State         string
	Hostname      string
	EnvironmentID string
}

// Environment is a representation of an Engine Yard Cloud environment.
type Environment struct {
	ID   string
	Name string
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
