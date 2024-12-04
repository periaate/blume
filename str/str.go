// Package str provides type constraints and functions for string types.
package str

// import (
// 	"regexp"
// 	"strings"
//
// 	. "github.com/periaate/blume/gen"
// )
//
// // Contains returns a predicate that checks if the input string contains any of the given substrings.
// func Contains[S ~string](args ...S) Predicate[S] {
// 	return func(str S) bool {
// 		for _, s := range args {
// 			if strings.Contains(string(str), string(s)) {
// 				return true
// 			}
// 		}
// 		return false
// 	}
// }
//
// // HasPrefix returns a predicate that checks if the input string has any of the given prefixes.
// func HasPrefix[S ~string](args ...S) Predicate[S] {
// 	return func(str S) bool {
// 		l := Lim[S](len(str))(args)
// 		for _, arg := range l {
// 			if string(str[:len(arg)]) == string(arg) {
// 				return true
// 			}
// 		}
// 		return false
// 	}
// }
//
// // HasSuffix returns a predicate that checks if the input string has any of the given suffixes.
// func HasSuffix[S ~string](args ...S) Predicate[S] {
// 	return func(str S) bool {
// 		l := Lim[S](len(str))(args)
// 		for _, arg := range l {
// 			if string(str[len(str)-len(arg):]) == string(arg) {
// 				return true
// 			}
// 		}
// 		return false
// 	}
// }
//
// // ReplacePrefix replaces the prefix of a string if it matches any of the given patterns.
// func ReplacePrefix[S ~string](pats ...S) Transformer[S] {
// 	return func(str S) S {
// 		if len(pats)%2 != 0 {
// 			return str
// 		}
// 		for i := 0; i < len(pats); i += 2 {
// 			p := pats[i]
// 			if len(p) > len(str) {
// 				continue
// 			}
//
// 			if string(p) == string(str[:len(p)]) {
// 				return S(string(pats[i+1]) + string(str[len(p):]))
// 			}
// 		}
//
// 		return str
// 	}
// }
//
// // ReplaceSuffix replaces the suffix of a string if it matches any of the given patterns.
// func ReplaceSuffix[S ~string](pats ...S) Transformer[S] {
// 	return func(str S) S {
// 		if len(pats)%2 != 0 {
// 			return str
// 		}
// 		for i := 0; i < len(pats); i += 2 {
// 			if len(pats[i]) > len(str) {
// 				continue
// 			}
// 			p := pats[i]
//
// 			if string(p) == string(str[len(str)-len(p):]) {
// 				return S(string(str[:len(str)-len(p)]) + string(pats[i+1]))
// 			}
// 		}
//
// 		return str
// 	}
// }
//
// // Replace replaces any found substrings with the patterns given.
// func Replace[S ~string](pats ...S) Transformer[S] {
// 	return func(str S) S {
// 		if len(pats)%2 != 0 {
// 			return str
// 		}
// 		for i := 0; i < len(pats); i += 2 {
// 			str = S(strings.ReplaceAll(string(str), string(pats[i]), string(pats[i+1])))
// 		}
// 		return str
// 	}
// }
//
// // ReplaceRegex replaces substrings matching a regex pattern.
// func ReplaceRegex[S ~string](pat string, rep string) Transformer[S] {
// 	matcher, err := regexp.Compile(pat)
// 	if err != nil {
// 		return func(_ S) (_ S) { return }
// 	}
// 	return func(s S) S {
// 		return S(matcher.ReplaceAll([]byte(string(s)), []byte(rep)))
// 	}
// }
//
// // Shift removes the first `count` characters from a string.
// func Shift[S ~string](count int) Transformer[S] {
// 	return func(a S) (res S) {
// 		if len(a) < count {
// 			return
// 		}
//
// 		return S(string(a[count:]))
// 	}
// }
//
// // Pop removes all but the first `count` characters from a string.
// func Pop[S ~string](count int) Transformer[S] {
// 	return func(a S) (res S) {
// 		if len(a) < count {
// 			return
// 		}
//
// 		return S(string(a[:count]))
// 	}
// }
