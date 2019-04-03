package bash

import (
	"os/exec"
	"syscall"
)

type ExecService struct{}

func NewExecService() *ExecService {
	return &ExecService{}
}

func (service *ExecService) Run(command string) int {
	return Run(command)
}

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
