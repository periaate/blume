package blume

import (
	"fmt"
	"os"
)

// Exit the program with a console log
func (f String) Exit(args ...any) { Exit(args...) }

// Exit the program with a console log
func Exit(args ...any) { fmt.Printf("%s\n", ToArray(args).Join(" ")); os.Exit(1) }

func (f String) ExitWith(n int) { fmt.Print(f); os.Exit(n) }
func ExitWith(n int, args ...any) { fmt.Printf("%s", ToArray(args).Join(" ")); os.Exit(n) }

func ExitsWith[A any](n int) func(arg A) A { return func(arg A) A { ExitWith(n, arg); return arg } }

func OrExit[A, B any](either Either[A, B], args ...any) (res A) {
	if !either.IsOk() {
		Exit(P.Printf("%s [%v]", P.S(args...), either.Other))
	}
	return either.Value
}
