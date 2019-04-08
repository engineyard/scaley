package bash

import (
	"os/exec"
	"syscall"
)

// ExecService is a service that knows how to execute external commands via
// bash.
type ExecService struct{}

// NewExecService creates and returns a new ExecService instance.
func NewExecService() *ExecService {
	return &ExecService{}
}

// Run takes a command and returns the result of running it with bash.Run.
func (service *ExecService) Run(command string) int {
	return Run(command)
}

// Run takes a command and returns the system status code that results from
// running said command via bash.
var Run = func(command string) int {
	cmd := exec.Command("bash", "-c", command)
	var waitStatus syscall.WaitStatus

	if err := cmd.Run(); err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			waitStatus = exitError.Sys().(syscall.WaitStatus)
			return waitStatus.ExitStatus()
		}
	}

	return 0
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
