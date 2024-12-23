package blume

import (
	"regexp"
	"sort"
	"strings"
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

// Contains returns a predicate that checks if the input string contains any of the given substrings.
func Contains[S ~string](args ...S) func(S) bool {
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
func HasPrefix[S ~string](args ...S) func(S) bool {
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
func HasSuffix[S ~string](args ...S) func(S) bool {
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

// ReplacePrefix replaces the prefix of a string if it matches any of the given patterns.
func ReplacePrefix[S ~string](pats ...S) func(S) S {
	return func(str S) S {
		if len(pats)%2 != 0 {
			return str
		}
		for i := 0; i < len(pats); i += 2 {
			p := pats[i]
			if len(p) > len(str) {
				continue
			}

			// blumefmt incorrectly inlines this
			if string(p) == string(str[:len(p)]) {
				return S(string(pats[i+1]) + string(str[len(p):]))
			}
		}

		return str
	}
}

// ReplaceSuffix replaces the suffix of a string if it matches any of the given patterns.
func ReplaceSuffix[S ~string](pats ...S) func(S) S {
	return func(str S) S {
		if len(pats)%2 != 0 {
			return str
		}
		for i := 0; i < len(pats); i += 2 {
			if len(pats[i]) > len(str) {
				continue
			}
			p := pats[i]

			if string(p) == string(str[len(str)-len(p):]) {
				// blumefmt incorrectly inlines this
				a := string(str[:len(str)-len(p)])
				b := string(pats[i+1])
				return S(a + b)
			}
		}

		return str
	}
}

// Replace replaces any found substrings with the patterns given.
func Replace[S ~string](pats ...S) func(S) S {
	return func(str S) S {
		if len(pats)%2 != 0 {
			return str
		}
		for i := 0; i < len(pats); i += 2 {
			str = S(strings.ReplaceAll(string(str), string(pats[i]), string(pats[i+1])))
		}
		return str
	}
}

// ReplaceRegex replaces substrings matching a regex pattern.
func MatchRegex[S ~string](pats ...S) func(S) bool {
	funcs := make([]func(S) bool, len(pats))
	for i, pat := range pats {
		matcher, err := regexp.Compile(string(pat))
		if err != nil {
			return func(_ S) (_ bool) { return }
		}
		funcs[i] = func(s S) bool { return matcher.MatchString(string(s)) }
	}
	return PredOr(funcs...)
}

// ReplaceRegex replaces substrings matching a regex pattern.
func ReplaceRegex[S ~string](pat string, rep string) func(S) S {
	matcher, err := regexp.Compile(pat)
	if err != nil {
		return func(_ S) (_ S) { return }
	}
	return func(s S) S { return S(matcher.ReplaceAll([]byte(string(s)), []byte(rep))) }
}

// Shift removes the first `count` characters from a string.
func Shift[S ~string](count int) func(S) S {
	return func(a S) (res S) {
		if len(a) < count {
			return
		}
		return S(string(a[count:]))
	}
}

// Pop removes all but the first `count` characters from a string.
func Pop[S ~string](count int) func(S) S {
	return func(a S) (res S) {
		if len(a) < count {
			return
		}
		return S(string(a[:count]))
	}
}

func Split(str string, keep bool, match ...string) (res []string) {
	if len(match) == 0 || len(str) == 0 {
		return []string{str}
	}

	sort.SliceStable(match, func(i, j int) bool {
		return len(match[i]) > len(match[j])
	})

	var lastI int
	for i := 0; i < len(str); i++ {
		for _, pattern := range match {
			switch {
			case i+len(pattern) > len(str):
				continue
			case str[i:i+len(pattern)] != pattern:
				continue
			case len(str[lastI:i]) != 0:
				res = append(res, str[lastI:i])
			}

			lastI = i + len(pattern)
			if len(pattern) != 0 {
				if keep {
					res = append(res, str[i:len(pattern)+i])
				}
				i += len(pattern) - 1
			}
			break
		}
	}

	if len(str[lastI:]) != 0 {
		res = append(res, str[lastI:])
	}

	return res
}
