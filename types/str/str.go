package str

import (
	"fmt"
	"sort"
	"strings"

	"regexp"

	"github.com/periaate/blume/pred"
	"github.com/periaate/blume/pred/has"
)

// ReplacePrefix replaces the prefix of a string if it matches any of the given patterns.
func ReplacePrefix(str string, pats ...string) (string, error) {
	if len(pats)%2 != 0 {
		return "", fmt.Errorf("patterns must be divisible by paired: [pattern, replacement]")
	}
	for i := 0; i < len(pats); i += 2 {
		p := pats[i]
		if len(p) > len(str) {
			continue
		}
		if p == str[:len(p)] {
			return pats[i+1] + str[len(p):], nil
		}
	}
	return str, nil
}

// Replacestringuffix replaces the suffix of a string if it matches any of the given patterns.
func ReplaceSuffix(str string, pats ...string) (string, error) {
	if len(pats)%2 != 0 {
		return "", fmt.Errorf("patterns must be divisible by paired: [pattern, replacement]")
	}
	for i := 0; i < len(pats); i += 2 {
		if len(pats[i]) > len(str) {
			continue
		}
		p := pats[i]
		if p == str[len(str)-len(p):] {
			return str[:len(str)-len(p)] + pats[i+1], nil
		}
	}
	return str, nil
}

// Replace replaces any found substrings with the patterns given.
func ReplaceAll(str string, pats ...string) (string, error) {
	if len(pats)%2 != 0 {
		return "", fmt.Errorf("patterns must be divisible by paired: [pattern, replacement]")
	}
	for i := 0; i < len(pats); i += 2 {
		str = strings.ReplaceAll(str, pats[i], pats[i+1])
	}
	return str, nil
}

// ReplaceRegex replaces substrings matching a regex pattern.
func ReplaceRegex(str string, pats ...string) (string, error) {
	if len(pats)%2 != 0 {
		return "", fmt.Errorf("patterns must be divisible by paired: [pattern, replacement]")
	}
	for i := 0; i < len(pats); i += 2 {
		pat, rep := pats[i], pats[i+1]
		matcher, err := regexp.Compile(pat)
		if err != nil {
			return "", err
		}
		str = matcher.ReplaceAllString(str, rep)
	}
	return str, nil
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

func TrimPrefixes(pats ...string) func(string) string {
	return func(inp string) string {
		for _, pat := range pats {
			if has.Prefix(pat)(inp) {
				return string(strings.TrimPrefix(inp, pat))
			}
		}
		return inp
	}
}

func TrimSuffixes(pats ...string) func(string) string {
	return func(inp string) string {
		for _, pat := range pats {
			if has.Suffix(pat)(inp) {
				return string(strings.TrimSuffix(inp, pat))
			}
		}
		return inp
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
	return pred.Or(funcs...)
}
