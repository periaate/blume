package filter

import (
	"github.com/periaate/blume/pred"
)

func Filter[A any](pred func(A) bool) func([]A) []A {
	return func(args []A) (res []A) {
		res = make([]A, 0, len(args))
		for _, arg := range args {
			if pred(arg) {
				res = append(res, arg)
			}
		}
		return res
	}
}

// Any filters an array such that elements pass at least one predicate.
func Any[A any](a ...func(A) bool) func([]A) []A { return Filter(pred.Or(a...)) }

// Every filters an array such that element pass all predicates.
func Every[A any](a ...func(A) bool) func([]A) []A { return Filter(pred.And(a...)) }
