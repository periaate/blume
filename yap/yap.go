package yap

import (
	"fmt"
	"os"
)

func Info(msg string, args ...any) {
	fmt.Fprintf(
		os.Stdout,
		"%v %s",
		// String("I"),
	)
}

func Debug(msg string, args ...any) {}

func Error(msg string, args ...any) {}
