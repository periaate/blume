//go:build !windows
// +build !windows

package system

import (
	"os/exec"
	"syscall"
)

// Stop stops the command and all its child processes.
func Stop(cmd *exec.Cmd) {
	// https://stackoverflow.com/questions/22470193/why-wont-go-kill-a-child-process-correctly
	// https://medium.com/@felixge/killing-a-child-process-and-all-of-its-children-in-go-54079af94773
	pgid := -cmd.Process.Pid
	_ = syscall.Kill(pgid, syscall.SIGTERM)
}
