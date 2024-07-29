// Package str provides type constraints and functions for string types.
package str

import (
	"strings"

	"github.com/periaate/blume/core/gen"
)

// Contains returns a predicate that checks if the input string contains any of the given substrings.
func Contains[S ~string](args ...S) gen.Predicate[S] {
	return func(str S) bool {
		for _, s := range args {
			if strings.Contains(string(str), string(s)) {
				return true
			}
		}
		return false
	}
}

// HasPrefix returns a predicate that checks if the input string has any of the given prefixes.
func HasPrefix(args ...string) gen.Predicate[string] {
	return func(str string) bool {
		for _, arg := range Limit(str, args) {
			if str[:len(arg)] == arg {
				return true
			}
		}

		return false
	}
}

// HasSuffix returns a predicate that checks if the input string has any of the given suffixes.
func HasSuffix(args ...string) gen.Predicate[string] {
	return func(str string) bool {
		for _, arg := range Limit(str, args) {
			if str[len(str)-len(arg):] == arg {
				return true
			}
		}

		return false
	}
}

// Limit filters the args to be less than or equal to the given Max length.
func Limit[T ~string | ~[]any](Max T, args []T) (res []T) {
	for _, r := range args {
		if len(r) <= len(Max) {
			res = append(res, r)
		}
	}
	return
}

// Replace replaces any found sub strings with the patterns given.
// Must have an even number of patterns. {from, to}
// Replacement done in the given order.
func Replace(pats ...string) gen.Monadic[string, string] {
	return func(str string) string {
		if len(pats)%2 != 0 {
			return str
		}
		for i := 0; i < len(pats); i += 2 {
			str = strings.ReplaceAll(str, pats[i], pats[i+1])
		}
		return str
	}
}

// func IsDigit(str string) bool {
// 	for _, r := range str {
// 		switch r {
// 		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
// 			continue
// 		default:
// 			return false
// 		}
// 	}
// 	return true
// }
//
// func IsNumber(str string) bool {
// 	if HasPrefix("-", "+")(str) {
// 		str = Slice(0, 1, str)
// 	}
// 	for _, r := range str {
// 		switch r {
// 		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.', ',':
// 			continue
// 		default:
// 			return false
// 		}
// 	}
// 	return true
// }
