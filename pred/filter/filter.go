package filter

import (
	"github.com/periaate/blume/pred"
)

// Any filters an array such that elements pass at least one predicate.
func Any[A any](a ...func(A) bool) func([]A) []A { return pred.Filter(pred.Or(a...)) }

// Every filters an array such that element pass all predicates.
func Every[A any](a ...func(A) bool) func([]A) []A { return pred.Filter(pred.And(a...)) }
