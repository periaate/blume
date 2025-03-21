package blume

import (
	"regexp"
	"sort"
	"strings"
)

// All returns true if all arguments pass the [Predicate].
func All[A any](fns ...Pred[A]) Pred[[]A] {
	fn := PredAnd(fns...)
	return func(args []A) bool {
		for _, arg := range args {
			if !fn(arg) {
				return false
			}
		}
		return true
	}
}

// Any returns true if any argument passes the [Predicate].
func Any[A any](fns ...Pred[A]) Pred[[]A] {
	fn := PredOr(fns...)
	return func(args []A) bool {
		for _, arg := range args {
			if fn(arg) {
				return true
			}
		}
		return false
	}
}

// Filter returns a slice of arguments that pass the [Predicate].
func Filter[A any](fns ...Pred[A]) func([]A) []A {
	fn := PredAnd(fns...)
	return func(args []A) (res []A) {
		res = make([]A, 0, len(args))
		for _, arg := range args {
			if fn(arg) {
				res = append(res, arg)
			}
		}
		return res
	}
}

// Map applies the function to each argument and returns the results.
func Map[A, B any](fn func(A) B) func([]A) []B {
	return func(args []A) (res []B) {
		res = make([]B, 0, len(args))
		for _, arg := range args {
			res = append(res, fn(arg))
		}
		return res
	}
}

// Reduce applies the function to each argument and returns the result.
func Reduce[A any, B any](fn func(B, A) B, init B) func([]A) B {
	return func(args []A) B {
		res := init
		for _, arg := range args {
			res = fn(res, arg)
		}
		return res
	}
}

// Not negates a [Predicate].
func Not[A any](fn Pred[A]) Pred[A] { return func(t A) bool { return !fn(t) } }

// Is returns a [Predicate] that checks if the argument is in the list.
func Is[C comparable](A ...C) func(C) bool {
	in := make(map[C]bool, len(A))
	for _, a := range A {
		in[a] = true
	}
	return func(c C) bool {
		_, ok := in[c]
		return ok
	}
}

// First returns the first value which passes the [Predicate].
func First[A any](fns ...Pred[A]) func([]A) Option[A] {
	fn := PredOr[A](fns...)
	return func(args []A) Option[A] {
		for _, arg := range args {
			if fn(arg) {
				return Some(arg)
			}
		}
		return None[A]()
	}
}

func StrIs[A, B ~string](vals ...A) func(B) bool {
	is := Is(vals...)
	return func(b B) bool { return is(A(b)) }
}

// Pipe combines variadic [Transformer]s into a single [Transformer].
func Pipe[A any](fns ...func(A) A) func(A) A {
	return func(a A) A {
		for _, fn := range fns {
			a = fn(a)
		}
		return a
	}
}

// Pipe combines variadic [Transformer]s into a single [Transformer].
func Pipes[A any](fns ...func(A) A) func([]A) []A { return Map(Pipe(fns...)) }

// Cat concatenates two [FnA] functions into a single [FnA] function.
func Cat[A, B, C any](a func(A) B, b func(B) C) func(A) C { return func(c A) C { return b(a(c)) } }

func PredAnd[A any](preds ...Pred[A]) Pred[A] {
	return func(a A) bool {
		for _, pred := range preds {
			if !pred(a) {
				return false
			}
		}
		return true
	}
}

func PredOr[A any](preds ...Pred[A]) Pred[A] {
	return func(a A) bool {
		for _, pred := range preds {
			if pred(a) {
				return true
			}
		}
		return false
	}
}

func Lazy[A any](fn func() A) func() A {
	var loaded bool
	var value A
	return func() A {
		if !loaded {
			value, loaded = fn(), true
		}
		return value
	}
}

func Memo[K comparable, V any](fn func(K) V) func(K) V {
	values := map[K]V{}
	return func(input K) (res V) {
		res, ok := values[input]
		if !ok {
			res = fn(input)
			values[input] = res
		}
		return res
	}
}

func Negate[A any](fn Pred[A]) Pred[A] { return func(a A) bool { return !fn(a) } }

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

func Vals[K comparable, V any](m map[K]V) Array[V] {
	if m == nil {
		return Arr[V]()
	}
	arr := []V{}
	for _, v := range m {
		arr = append(arr, v)
	}
	return ToArray(arr)
}

func Keys[K comparable, V any](m map[K]V) Array[K] {
	if m == nil {
		return Arr[K]()
	}
	arr := []K{}
	for k := range m {
		arr = append(arr, k)
	}
	return ToArray(arr)
}
