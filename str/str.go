// Package str provides type constraints and functions for string types.
package str

import (
	"regexp"
	"strings"

	"github.com/periaate/blume/gen"
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
		l := gen.Lim[string](len(str))(args)
		for _, arg := range l {
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
		l := gen.Lim[string](len(str))(args)
		for _, arg := range l {
			if str[len(str)-len(arg):] == arg {
				return true
			}
		}

		return false
	}
}

// ReplacePrefix
func ReplacePrefix(pats ...string) gen.Transformer[string] {
	return func(str string) string {
		if len(pats)%2 != 0 {
			return str
		}
		for i := 0; i < len(pats); i += 2 {
			p := pats[i]
			if len(p) > len(str) {
				continue
			}

			if p == str[:len(p)] {
				return pats[i+1] + str[len(p):]
			}
		}

		return str
	}
}

// ReplaceSuffix
func ReplaceSuffix(pats ...string) gen.Transformer[string] {
	return func(str string) string {
		if len(pats)%2 != 0 {
			return str
		}
		for i := 0; i < len(pats); i += 2 {
			if len(pats[i]) > len(str) {
				continue
			}
			p := pats[i]

			if p == str[len(str)-len(p):] {
				return str[:len(str)-len(p)] + pats[i+1]
			}
		}

		return str
	}
}

// Replace replaces any found sub strings with the patterns given.
// Must have an even number of patterns. {from, to}
// Replacement done in the given order.
func Replace(pats ...string) gen.Transformer[string] {
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

func ReplaceRegex(pat string, rep string) gen.Transformer[string] {
	matcher, err := regexp.Compile(pat)
	if err != nil {
		return func(s string) string { return "" }
	}
	return func(s string) string {
		return string(matcher.ReplaceAll([]byte(s), []byte(rep)))
	}
}

func Shift(count int) gen.Transformer[string] {
	return func(a string) (res string) {
		if len(a) < count {
			return
		}

		return a[count:]
	}
}

func Pop(count int) gen.Transformer[string] {
	return func(a string) (res string) {
		if len(a) < count {
			return
		}

		return a[:count]
	}
}
