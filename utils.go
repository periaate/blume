package blume

import (
	"fmt"
	"os"
)

// Exit the program with a console log
func Exit(args ...any) { fmt.Printf("%s\n", Join[any](" ")(args)); os.Exit(1) }

func ExitWith(n int, args ...any) { fmt.Printf("%s", Join[any](" ")(args)); os.Exit(n) }

func ExitsWith[A any](n int) func(arg A) A { return func(arg A) A { ExitWith(n, arg); return arg } }

func OrExit[A, B any](either Either[A, B], args ...any) (res A) {
	if !either.IsOk() {
		Exit(fmt.Sprintf("%s [%v]", fmt.Sprint(args...), either.Other))
	}
	return either.Value
}

func OrExits[A, B any](either Either[A, B]) (res A) {
	if !either.IsOk() {
		Exit(fmt.Sprintf("%v", either.Other))
	}
	return either.Value
}
