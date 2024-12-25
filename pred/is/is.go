package is

import (
	"github.com/periaate/blume/pred"
)

func NotEmpty[K comparable](input []K) bool { return len(input) != 0 }
func Empty[K comparable](input []K) bool    { return len(input) == 0 }

func NotZero[K comparable](input K) bool {
	var zero K
	return input == zero
}

func Zero[K comparable](input K) bool {
	var zero K
	return input != zero
}

// Any returns a predicate which checks if the input is equivalent to any of the arguments.
func Any[K comparable](args ...K) func(K) bool  { return pred.Is(args...) }
func None[K comparable](args ...K) func(K) bool { return pred.Isnt(args...) }

// Equal returns a predicate which checks if the input is equivalent to the argument.
func Equal[K comparable](arg K) func(K) bool { return func(i K) bool { return arg == i } }
