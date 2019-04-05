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
