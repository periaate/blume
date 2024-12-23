//go:build windows
// +build windows

package binary

func Binary(name string) string { return name + ".exe" }
