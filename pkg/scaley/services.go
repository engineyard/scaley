package scaley

// Services is a handy collection of the various services one uses to perform
// tasks within Scaley.
type Services struct {
	Groups       GroupService
	Servers      ServerService
	Environments EnvironmentService
	Scripts      ScalingScriptService
	Locker       LockService
	Runner       ExecService
	Log          LogService
}

// GroupService is an interface that describes an object that knows how to
// interact with a Group.
type GroupService interface {
	Get(string) (*Group, error)
}

// ServerService is an interface that describes an object that knows how to
// interact with a Server.
type ServerService interface {
	Get(string) (*Server, error)
	Start(*Server) error
	Stop(*Server) error
}

// EnvironmentService is an interface that describes an object that knows how
// to interact with an Environment.
type EnvironmentService interface {
	Get(string) (*Environment, error)
	Configure(*Environment) error
}

// LockService is an interface that describes an object that knows how to
// deal with Group locking.
type LockService interface {
	Lock(*Group) error
	Unlock(*Group) error
	Locked(*Group) bool
}

// ExecService is an interface that describes an object that knows how to
// execute external commands.
type ExecService interface {
	Run(string) int
}

// ScalingScriptService is an interface that describes an object that knows how
// to interct with a ScalingScript.
type ScalingScriptService interface {
	Exists(string) bool
}

// LogService is an interface that describes an object that provides logging
// capabilities.
type LogService interface {
	Info(*Group, string)
	Failure(*Group, string)
	Success(*Group, string)
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
