package yap

import "testing"

func TestInfo(t *testing.T) {
	SetLevel(L_Debug)
	Error("This is an error message", "meow", 123)
	Info("This is an info message", "meow", 123)
	Debug("This is a debug message", "meow", 123)
}
