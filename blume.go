package blume

import (
	"regexp"
	"slices"
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
		return slices.ContainsFunc(args, fn)
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

func AMap[A, B any](fn func(A) B) func(Array[A]) Array[B] {
	return func(args Array[A]) Array[B] {
		res := make([]B, 0, len(args.Value))
		for _, arg := range args.Value {
			res = append(res, fn(arg))
		}
		return ToArray(res)
	}
}

func Maps[A, B any](fn func(A) B) func([]A) Array[B] { return Cat(Map(fn), ToArray) }

func FlatMap[A, B any](fn func(A) []B) func([]A) []B {
	return func(args []A) (res []B) {
		for _, arg := range args {
			res = append(res, fn(arg)...)
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

func Fold[A any, B any](fn func(B, A) B, init ...B) func([]A) B {
	var in B
	if len(init) > 0 { in = init[0] }
	return func(args []A) B {
		res := in
		for _, arg := range args {
			res = fn(res, arg)
		}
		return res
	}
}

// Not negates a [Predicate].
func Not[A any](fn Pred[A]) Pred[A] { return func(t A) bool { return !fn(t) } }

func IsZero[K comparable](a K) bool {
	var def K
	return a == def
}

func IsEvery[C comparable](args ...C) func(C) bool { return PredAnd(Map(V2M[C](Is))(args)...) }

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
	fn := PredOr(fns...)
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
func Catn[A, B any](a func(A) B, b func(B)) func(A) { return func(c A) { b(a(c)) } }
func Catc[A, B any](a func() A, b func(A) B) func() B { return func() B { return b(a()) } }

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

func LazyW[A, B any](fn func(A) B, input A) func() B {
	var loaded bool
	var value B
	return func() B {
		if !loaded { value, loaded = fn(input), true }
		return value
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

func Limit[A ~string | ~[]any](Max int) func([]A) []A {
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
func Contains(args ...S) func(S) bool {
	return func(str S) bool {
		for _, s := range args {
			if strings.Contains(string(str), string(s)) {
				return true
			}
		}
		return false
	}
}

func Includes[K comparable](inclusive bool) func(args ...K) func([]K) bool {
	return func(args ...K) func([]K) bool {
		var pred Pred[K]
		if inclusive { pred = Is(args...) } else { pred = IsEvery(args...) }
		return func(arr []K) bool { return slices.ContainsFunc(arr, pred) }
	}
}

// HasPrefix returns a predicate that checks if the input string has any of the given prefixes.
func HasPrefix(args ...S) func(S) bool {
	return func(str S) bool {
		l := Limit[S](len(str))(args)
		for _, arg := range l {
			if string(str[:len(arg)]) == string(arg) {
				return true
			}
		}
		return false
	}
}

// HasSuffix returns a predicate that checks if the input string has any of the given suffixes.
func HasSuffix(args ...S) func(S) bool {
	return func(str S) bool {
		l := Limit[S](len(str))(args)
		for _, arg := range l {
			if string(str[len(str)-len(arg):]) == string(arg) {
				return true
			}
		}
		return false
	}
}

// ReplacePrefix replaces the prefix of a string if it matches any of the given patterns.
func ReplacePrefix(pats ...S) func(S) S {
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
func ReplaceSuffix(pats ...S) func(S) S {
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
func Replace(pats ...S) func(S) S {
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
func MatchRegex(pats ...S) func(S) bool {
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
func ReplaceRegex(pat String, rep String) func(S) S {
	matcher, err := regexp.Compile(pat.String())
	if err != nil {
		return func(_ S) (_ S) { return }
	}
	return func(s S) S { return S(matcher.ReplaceAll([]byte(string(s)), []byte(rep))) }
}

// Shift removes the first `count` characters from a string.
func Shift(count int) func(S) S {
	return func(a S) (res S) {
		if len(a) < count {
			return
		}
		return S(string(a[count:]))
	}
}

// Pop removes all but the first `count` characters from a string.
func Pop(count int) func(S) S {
	return func(a S) (res S) {
		if len(a) < count {
			return
		}
		return S(string(a[:count]))
	}
}

func Splits(keep bool, match ...String) func(String) []String {
	return func(str String) (res []S) {
		if len(match) == 0 || len(str) == 0 {
			return []S{str}
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
}

func Split(str String, keep bool, match ...String) (res []String) {
	if len(match) == 0 || len(str) == 0 {
		return []String{str}
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
