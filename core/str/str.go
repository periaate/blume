package str

import (
	"strings"

	"github.com/periaate/blume/core/gen"
	"github.com/periaate/blume/core/num"
)

func Contains(args ...string) gen.Predicate[string] {
	return func(str string) bool {
		for _, s := range args {
			if strings.Contains(str, s) {
				return true
			}
		}
		return false
	}
}

func HasPrefix(args ...string) gen.Predicate[string] {
	return func(str string) bool {
		for _, arg := range Limit(str, args...) {
			if str[:len(arg)] == arg {
				return true
			}
		}

		return false
	}
}

func HasSuffix(args ...string) gen.Predicate[string] {
	return func(str string) bool {
		for _, arg := range Limit(str, args...) {
			if str[len(str)-len(arg):] == arg {
				return true
			}
		}

		return false
	}
}

func Limit[T ~string | ~[]T](i T, ts ...T) (res []T) {
	for _, r := range ts {
		if len(r) <= len(i) {
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

func IsDigit(str string) bool {
	for _, r := range str {
		switch r {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			continue
		default:
			return false
		}
	}
	return true
}

func Slice(from, to int, inp string) (res string) {
	if from == to {
		return
	}
	l := len(inp)
	from = num.SmartClamp(from, l)
	to = num.SmartClamp(to, l)

	if from > to {
		return inp[to:from]
	}

	return inp[from:to]
}

func IsNumber(str string) bool {
	if HasPrefix("-", "+")(str) {
		str = Slice(0, 1, str)
	}
	for _, r := range str {
		switch r {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.', ',':
			continue
		default:
			return false
		}
	}
	return true
}
