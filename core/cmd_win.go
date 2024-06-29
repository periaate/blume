//go:build windows
// +build windows

package core

import (
	"os/exec"
	"strconv"
)

// Stop stops the command and all its child processes.
func Stop(cmd *exec.Cmd) {
	// https://stackoverflow.com/a/44551450
	killCmd := exec.Command("taskkill.exe", "/t", "/f", "/pid", strconv.Itoa(cmd.Process.Pid))
	_ = killCmd.Run()
}
