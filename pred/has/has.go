package has

import (
	"strings"

	"github.com/periaate/blume/pred"
)

func limit[A ~string | ~[]any](Max int) func([]A) []A {
	return func(args []A) (res []A) {
		for _, a := range args {
			if len(a) <= Max {
				res = append(res, a)
			}
		}
		return
	}
}

func None(args ...string) func(string) bool     { return pred.Not(Any(args...)) }
func NotEvery(args ...string) func(string) bool { return pred.Not(Every(args...)) }

func Every(args ...string) func(string) bool {
	return func(str string) bool {
		for _, s := range args {
			if !strings.Contains(string(str), string(s)) {
				return false
			}
		}
		return true
	}
}

// Contains returns a predicate that checks if the input string contains any of the given substrings.
func Any(args ...string) func(string) bool {
	return func(str string) bool {
		for _, s := range args {
			if strings.Contains(string(str), string(s)) {
				return true
			}
		}
		return false
	}
}

// HasPrefix returns a predicate that checks if the input string has any of the given prefixes.
func Prefix[S ~string](args ...S) func(S) bool {
	return func(str S) bool {
		l := limit[S](len(str))(args)
		for _, arg := range l {
			if string(str[:len(arg)]) == string(arg) {
				return true
			}
		}
		return false
	}
}

// HasSuffix returns a predicate that checks if the input string has any of the given suffixes.
func Suffix[S ~string](args ...S) func(S) bool {
	return func(str S) bool {
		l := limit[S](len(str))(args)
		for _, arg := range l {
			if string(str[len(str)-len(arg):]) == string(arg) {
				return true
			}
		}
		return false
	}
}
