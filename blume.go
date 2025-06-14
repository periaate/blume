package blume

import (
	"fmt"
	"reflect"
	"regexp"
	"slices"
	"sort"
	"strings"
)

// All returns true if all arguments pass the [Predicate].
func All[T any](fns ...Pred[T]) Pred[[]T] {
	fn := PredAnd(fns...)
	return func(args []T) bool {
		for _, arg := range args {
			if !fn(arg) {
				return false
			}
		}
		return true
	}
}

// Filter returns a slice of arguments that pass the [Predicate].
func Filter[T any](fns ...Pred[T]) func(Array[T]) Array[T] {
	fn := PredAnd(fns...)
	return func(args Array[T]) (res Array[T]) {
		for _, arg := range args {
			if fn(arg) {
				res = append(res, arg)
			}
		}
		return res
	}
}

func FilterMap[I, O any](fn func(I) Option[O]) func(A[I]) A[O] {
	return func(arr A[I]) A[O] {
		res := []O{}
		for _, val := range arr {
			if val := fn(val); val.IsOk() {
				res = append(res, val.Value)
			}
		}
		return res
	}
}

type MapFn[I, O any] interface { Mapper[I, O] | TTVar[I, O] | AVar[O] | Say | TVar[I, O] }
type Flatter[T1, T2 any] = func(T1) Option[T2]
type Mapper[T1, T2 any]  = func(T1) T2
type TTVar[T1, T2 any]   = func(...T1) T2
type AVar[T any]         = func(...any) T
type Say                 = func(...any)
type TVar[T1, T2 any]    = func(T1, ...any) T2
type Shout[T any]        = func(T)
type Pred[T any]         = func(T) bool

func Over[I, O any, Fn Flatter[I, O] | Mapper[I, O] | TTVar[I, O] | AVar[O] | Say | TVar[I, O] | Shout[I] | Pred[I]](arg Fn) (res func(Array[I]) Array[O]) {
	switch fn := any(arg).(type) {
	case Say     : return As(res, Each(func(t I) { fn(t) }))
	case Shout[I]: return As(res, Each(fn))

	case Pred   [I]   : return As(res, Filter(fn))
	case Flatter[I, O]: return FilterMap(fn)

	case TVar  [I, O]: return Map[I, O](func(t I) O { return fn(t) })
	case AVar  [O]   : return Map[I, O](func(t I) O { return fn(t) })
	case TTVar [I, O]: return Map [I, O](func(t I) O { return fn(t) })
	case Mapper[I, O]: return Map  [I, O](fn)
	default          : return
	}
}

func OverCoax[I, O any](arg any) (res func(Array[I]) Array[O]) {
	switch fn := arg.(type) {
	case Say           : return As(res, Each(func(t I) { fn(t) }))
	case Shout[I]      : return As(res, Each(fn))

	case Pred[I]       : return As(res, Filter(fn))
	case Flatter[I, O] : return FilterMap(fn)

	case TVar[I, O]    : return Map[I, O](func(t I) O { return fn(t) })
	case AVar[O]       : return Map[I, O](func(t I) O { return fn(t) })
	case TTVar[I, O]   : return Map[I, O](func(t I) O { return fn(t) })
	case Mapper[I, O]  : return Map[I, O](fn)
	default            : return
	}
}

func As[target any](_ target, arg any) target {
	fn, ok := arg.(target)
	if !ok { panic("as called wit invalid function") }
	return fn
}

func Each[T any, Arr Array[T]](fn func(T)) func(Arr) Arr {
	return func(arr Arr) Arr {
		for _, value := range arr {
			fn(value)
		}
		return arr
	}
}


// Map applies the function to each argument and returns the results.
func Map[I, O any, Fn MapFn[I, O]](arg Fn) func(Array[I]) Array[O] {
	var fn func(I) O
	switch fun := any(arg).(type) {
	case TVar[I, O]    : fn = func(t I) O { return fun(t) }
	case AVar[O]       : fn = func(t I) O { return fun(t) }
	case TTVar[I, O]   : fn = func(t I) O { return fun(t) }
	case Mapper[I, O]  : fn = fun }

	return func(args Array[I]) (res Array[O]) {
		res = make([]O, 0, len(args))
		for _, arg := range args { res = append(res, fn(arg)) }
		return res
	}
}

func FlatMap[I, O any](fn func(I) Array[O]) func(Array[I]) Array[O] {
	return func(args Array[I]) (res Array[O]) {
		for _, arg := range args {
			res = append(res, fn(arg)...)
		}
		return res
	}
}

// Reduce applies the function to each argument and returns the result.
func Reduce[T any, B any](fn func(B, T) B, init B) func([]T) B {
	return func(args []T) B {
		res := init
		for _, arg := range args {
			res = fn(res, arg)
		}
		return res
	}
}

func Fold[T any, B any](fn func(B, T) B, init ...B) func([]T) B {
	var in B
	if len(init) > 0 { in = init[0] }
	return func(args []T) B {
		res := in
		for _, arg := range args {
			res = fn(res, arg)
		}
		return res
	}
}

// Not negates a [Predicate].
func Not[T any](fn Pred[T]) Pred[T] { return func(t T) bool { return !fn(t) } }

func IsZero[K comparable](a K) bool {
	var def K
	return a == def
}

// Is returns a [Predicate] that checks if the argument is in the list.
func Is[C comparable](T ...C) func(C) bool {
	in := make(map[C]bool, len(T))
	for _, a := range T {
		in[a] = true
	}
	return func(c C) bool {
		_, ok := in[c]
		return ok
	}
}

// First returns the first value which passes the [Predicate].
func First[T any](fns ...Pred[T]) func([]T) Option[T] {
	fn := PredOr(fns...)
	return func(args []T) Option[T] {
		for _, arg := range args {
			if fn(arg) {
				return Some(arg)
			}
		}
		return None[T]()
	}
}

// Pipe runs a value through a pipeline or composes functions.
//
// If the first argument is a value, it executes a pipeline:
// T1, (T1) -> T2, (T2) -> T3, ..., (Tn-1) -> Tn
// and returns the final value: Tn
//
// If the first argument is a function, it composes a pipeline:
// (T1) -> T2, (T2) -> T3, ..., (Tn-1) -> Tn
// into a final function: (T1) -> Tn
func Pipe[Output any](values ...any) Output {
	// If no arguments are provided, return the zero value of the output type.
	var zero Output
	if len(values) == 0 {
		return zero
	}

	first := reflect.ValueOf(values[0])

	if first.Kind() != reflect.Func {
		if len(values) == 1 {
			// Only a single value was passed, return it.
			return first.Interface().(Output)
		}

		result := first
		for i := 1; i < len(values); i++ {
			fn := reflect.ValueOf(values[i])

			if fn.Kind() != reflect.Func {
				panic("Pipe Error: For value processing, all subsequent arguments must be functions.")
			}
			outputs := fn.Call([]reflect.Value{result})

			if len(outputs) == 0 {
				if i < len(values)-1 {
					panic("Pipe Error: A function in the middle of the pipeline returned no value, breaking the chain.")
				}
				return zero
			}
			result = outputs[0]
		}
		return result.Interface().(Output)
	}

	funcs := make([]reflect.Value, len(values))
	for i, v := range values {
		fn := reflect.ValueOf(v)
		if fn.Kind() != reflect.Func {
			panic("Pipe Error: For function composition, all arguments must be functions.")
		}
		funcs[i] = fn
	}

	for i := range len(funcs)-1 {
		if funcs[i].Type().NumOut() != funcs[i+1].Type().NumIn() {
			panic(fmt.Sprintf(
				"Pipe Error: Arity mismatch between function %d (returns %d values) and function %d (expects %d values).",
				i, funcs[i].Type().NumOut(), i+1, funcs[i+1].Type().NumIn(),
			))
		}

		for j := range funcs[i].Type().NumOut() {
			if funcs[i].Type().Out(j) != funcs[i+1].Type().In(j) {
				panic(fmt.Sprintf(
					"Pipe Error: Type mismatch between output %d of function %d (%s) and input %d of function %d (%s).",
					j, i, funcs[i].Type().Out(j), j, i+1, funcs[i+1].Type().In(j),
				))
			}
		}
	}

	firstFuncType := funcs[0].Type()
	lastFuncType := funcs[len(funcs)-1].Type()

	inTypes := make([]reflect.Type, firstFuncType.NumIn())
	for i := range firstFuncType.NumIn() {
		inTypes[i] = firstFuncType.In(i)
	}

	outTypes := make([]reflect.Type, lastFuncType.NumOut())
	for i := range lastFuncType.NumOut() {
		outTypes[i] = lastFuncType.Out(i)
	}

	composedFuncType := reflect.FuncOf(inTypes, outTypes, firstFuncType.IsVariadic())

	composedFuncImpl := func(args []reflect.Value) []reflect.Value {
		var currentResult = args
		for _, fn := range funcs {
			currentResult = fn.Call(currentResult)
		}
		return currentResult
	}

	composedFunc := reflect.MakeFunc(composedFuncType, composedFuncImpl)

	return composedFunc.Interface().(Output)
}

func Cat[T, B, C any](a func(T) B, b func(B) C) func(T) C { return func(c T) C { return b(a(c)) } }

func PredAnd[T any](preds ...Pred[T]) Pred[T] {
	return func(a T) bool {
		for _, pred := range preds {
			if !pred(a) {
				return false
			}
		}
		return true
	}
}

func PredOr[T any](preds ...Pred[T]) Pred[T] {
	return func(a T) bool {
		for _, pred := range preds {
			if pred(a) {
				return true
			}
		}
		return false
	}
}

func LazyW[T, B any](fn func(T) B, input T) func() B {
	var loaded bool
	var value B
	return func() B {
		if !loaded { value, loaded = fn(input), true }
		return value
	}
}

func Lazy[T any](fn func() T) func() T {
	var loaded bool
	var value T
	return func() T {
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

func Negate[T any](fn Pred[T]) Pred[T] { return func(a T) bool { return !fn(a) } }

func limit[T ~string | ~[]any](Max int) func([]T) []T {
	return func(args []T) (res []T) {
		for _, a := range args {
			if len(a) <= Max {
				res = append(res, a)
			}
		}
		return
	}
}

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

func Includes[K comparable](inclusive bool) func(args ...K) func([]K) bool {
	return func(args ...K) func([]K) bool {
		var pred Pred[K]
		// if inclusive { pred = Is(args...) } else { pred = IsEvery(args...) }
		pred = Is(args...)
		return func(arr []K) bool { return slices.ContainsFunc(arr, pred) }
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

func Vals[K comparable, V any](m map[K]V) (res Array[V]) {
	if m == nil { return }
	for _, v := range m {res = append(res, v)}
	return
}

func Keys[K comparable, V any](m map[K]V) (res Array[K]) {
	if m == nil { return }
	for k := range m {res = append(res, k)}
	return
}
