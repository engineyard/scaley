package basher

import (
	"os/exec"
	"syscall"
)

var Run func(command string) int

func init() {
	if Run == nil {
		Run = func(command string) int {
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
	}
}
