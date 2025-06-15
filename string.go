package blume

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Contains returns a predicate that checks if the input string contains any of the given substrings.
func Contains(args ...string) func(string) bool {
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
func HasPrefix(args ...string) func(string) bool {
	return func(str string) bool {
		l := limit[string](len(str))(args)
		for _, arg := range l {
			if string(str[:len(arg)]) == string(arg) {
				return true
			}
		}
		return false
	}
}

// HasSuffix returns a predicate that checks if the input string has any of the given suffixes.
func HasSuffix(args ...string) func(string) bool {
	return func(str string) bool {
		l := limit[string](len(str))(args)
		for _, arg := range l {
			if string(str[len(str)-len(arg):]) == string(arg) {
				return true
			}
		}
		return false
	}
}

// ReplacePrefix replaces the prefix of a string if it matches any of the given patterns.
func ReplacePrefix(pats ...string) func(string) string {
	return func(str string) string {
		if len(pats)%2 != 0 {
			return str
		}
		for i := 0; i < len(pats); i += 2 {
			p := pats[i]
			if len(p) > len(str) {
				continue
			}

			if string(p) == string(str[:len(p)]) {
				return string(string(pats[i+1]) + string(str[len(p):]))
			}
		}

		return str
	}
}

// ReplaceSuffix replaces the suffix of a string if it matches any of the given patterns.
func ReplaceSuffix(pats ...string) func(string) string {
	return func(str string) string {
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
				return string(a + b)
			}
		}

		return str
	}
}

// Replace replaces any found substrings with the patterns given.
func Replace(pats ...string) func(string) string {
	return func(str string) string {
		if len(pats)%2 != 0 {
			return str
		}
		for i := 0; i < len(pats); i += 2 {
			str = string(strings.ReplaceAll(string(str), string(pats[i]), string(pats[i+1])))
		}
		return str
	}
}

// ReplaceRegex replaces substrings matching a regex pattern.
func MatchRegex(pats ...string) func(string) bool {
	funcs := make([]func(string) bool, len(pats))
	for i, pat := range pats {
		matcher, err := regexp.Compile(string(pat))
		if err != nil {
			return func(_ string) (_ bool) { return }
		}
		funcs[i] = func(s string) bool { return matcher.Match([]byte(s)) }
	}
	return PredOr(funcs...)
}

// ReplaceRegex replaces substrings matching a regex pattern.
func ReplaceRegex(pat string, rep string) func(string) string {
	matcher, err := regexp.Compile(pat)
	if err != nil {
		return func(_ string) (_ string) { return }
	}
	return func(s string) string { return string(matcher.ReplaceAll([]byte(string(s)), []byte(rep))) }
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

func Join(sep string) func(args []any) string {
	return func(args []any) string {
		res := []string{}
		for _, arg := range args {
			res = append(res, fmt.Sprint(arg))
		}
		return strings.Join(res, sep)
	}
}

func EnsurePrefix(fix string) func(string) string {
	return func(s string) string {
		if HasPrefix(fix)(s) { return s }
		return fix + s
	}
}

func EnsureSuffix(fix string) func(string) string {
	return func(s string) string {
		if HasSuffix(fix)(s) { return s }
		return s + fix
	}
}

// SplitRegex keeps matches
func SplitRegex(pattern string) func(input string) []string {
	return func(input string) []string {
		re := regexp.MustCompile(pattern)
		indexes := re.FindAllStringIndex(input, -1)
		if len(indexes) == 0 {
			return []string{input}
		}

		result := make([]string, 0, 2*len(indexes)+1)
		lastEnd := 0

		for _, idx := range indexes {
			start, end := idx[0], idx[1]
			if start > lastEnd {
				result = append(result, input[lastEnd:start])
			}
			result = append(result, input[start:end])
			lastEnd = end
		}

		if lastEnd < len(input) {
			result = append(result, input[lastEnd:])
		}

		return result
	}
}

func ParseDuration(s string) (res time.Duration, ok bool) {
	var value time.Duration
	for _, v := range Split(s, false, " ") {
		var dur time.Duration
		dur, err := time.ParseDuration(v)
		if err != nil { return value, false }
		value += dur
	}

	return value, true
}

func ToInt[S ~string](s S) Option[int] {
	i, err := strconv.Atoi(string(s))
	return Either[int, bool]{Value: int(i), Other: err == nil}
}
func ToInt8[S ~string](s S) Option[int8] {
	i, err := strconv.ParseInt(string(s), 10, 8)
	return Either[int8, bool]{Value: int8(i), Other: err == nil}
}
func ToInt16[S ~string](s S) Option[int16] {
	i, err := strconv.ParseInt(string(s), 10, 16)
	return Either[int16, bool]{Value: int16(i), Other: err == nil}
}
func ToInt32[S ~string](s S) Option[int32] {
	i, err := strconv.ParseInt(string(s), 10, 32)
	return Either[int32, bool]{Value: int32(i), Other: err == nil}
}
func ToInt64[S ~string](s S) Option[int64] {
	i, err := strconv.ParseInt(string(s), 10, 64)
	return Either[int64, bool]{Value: int64(i), Other: err == nil}
}
func ToUint[S ~string](s S) Option[uint] {
	i, err := strconv.ParseUint(string(s), 10, 0)
	return Either[uint, bool]{Value: uint(i), Other: err == nil}
}
func ToUint8[S ~string](s S) Option[uint8] {
	i, err := strconv.ParseUint(string(s), 10, 8)
	return Either[uint8, bool]{Value: uint8(i), Other: err == nil}
}
func ToUint16[S ~string](s S) Option[uint16] {
	i, err := strconv.ParseUint(string(s), 10, 16)
	return Either[uint16, bool]{Value: uint16(i), Other: err == nil}
}
func ToUint32[S ~string](s S) Option[uint32] {
	i, err := strconv.ParseUint(string(s), 10, 32)
	return Either[uint32, bool]{Value: uint32(i), Other: err == nil}
}
func ToUint64[S ~string](s S) Option[uint64] {
	i, err := strconv.ParseUint(string(s), 10, 64)
	return Either[uint64, bool]{Value: uint64(i), Other: err == nil}
}
func ToFloat32[S ~string](s S) Option[float32] {
	i, err := strconv.ParseFloat(string(s), 32)
	return Either[float32, bool]{Value: float32(i), Other: err == nil}
}
func ToFloat64[S ~string](s S) Option[float64] {
	i, err := strconv.ParseFloat(string(s), 64)
	return Either[float64, bool]{Value: float64(i), Other: err == nil}
}
