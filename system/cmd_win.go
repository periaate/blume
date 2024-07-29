//go:build windows
// +build windows

package system

import (
	"os/exec"
	"strconv"
)

// Stop stops the command and all its child processes.
func Stop(cmd *exec.Cmd) error {
	// https://stackoverflow.com/a/44551450
	killCmd := exec.Command("taskkill.exe", "/t", "/f", "/pid", strconv.Itoa(cmd.Process.Pid))
	return killCmd.Run()
}
