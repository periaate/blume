package str

import (
	"sort"
	"strings"
)

func Every(anyM bool, fn func(string, ...string) bool, tars []string, pats ...string) bool {
	for _, str := range tars {
		if fn(str, pats...) == anyM {
			return anyM
		}
	}
	return !anyM
}

func Some(fn func(string, ...string) bool, tars []string, pats ...string) bool {
	for _, str := range tars {
		if fn(str, pats...) {
			return true
		}
	}
	return false
}

func Contains(str string, sub ...string) bool {
	sub = Limit(str, sub...)
	sort.Slice(sub, func(i, j int) bool {
		return len(sub[i]) < len(sub[j])
	})

	var lens []int
	for _, v := range sub {
		lens = append(lens, len(v))
	}

	li := 0
	l := lens[li]
	for i := range str {
		if i+l > len(str) {
			li++
			if li >= len(lens) {
				return false
			}

			l = lens[li]
		}
		for _, s := range sub[li:] {
			if strings.Contains(str, s) {
				return true
			}
		}
	}

	return false
}

func HasPrefix(str string, pre ...string) bool {
	for _, s := range Limit(str, pre...) {
		if str[:len(s)] == s {
			return true
		}
	}

	return false
}

func HasSuffix(str string, suf ...string) bool {
	for _, s := range Limit(str, suf...) {
		if str[len(str)-len(s):] == s {
			return true
		}
	}

	return false
}

func Limit[T ~string | ~[]T](i T, ts ...T) (res []T) {
	for _, r := range ts {
		if len(r) <= len(i) {
			res = append(res, r)
		}
	}
	return
}

func Is(str string, is ...string) bool {
	is = Limit(str, is...)
	for _, s := range is {
		if str == s {
			return true
		}
	}

	return false
}

// ReplaceAny replaces any found sub strings with the patterns given.
// Must have an even number of patterns. {from, to}
func ReplaceAny(str string, pats ...string) string {
	if len(pats)%2 != 0 {
		return str
	}
	for i := 0; i < len(pats); i += 2 {
		str = strings.ReplaceAll(str, pats[i], pats[i+1])
	}
	return str
}

func IsDigit(str ...string) bool {
	for _, s := range str {
		for _, r := range s {
			switch r {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				continue
			default:
				return false
			}
		}
	}
	return true
}
