package blume

import (
	"fmt"
	"reflect"
	"slices"
)

type Value struct {
	reflect.Value
	Then *Value
}

func (val *Value) IsFunc  () bool { return val.Kind() == reflect.Func   }
func (val *Value) IsBool  () bool { return val.Kind() == reflect.Bool   }
func (val *Value) IsInt   () bool { return val.Kind() == reflect.Int    }
func (val *Value) IsUint  () bool { return val.Kind() == reflect.Uint   }
func (val *Value) IsUint8 () bool { return val.Kind() == reflect.Uint8  }
func (val *Value) IsUint16() bool { return val.Kind() == reflect.Uint16 }
func (val *Value) IsUint32() bool { return val.Kind() == reflect.Uint32 }
func (val *Value) IsUint64() bool { return val.Kind() == reflect.Uint64 }
func (val *Value) IsInt8  () bool { return val.Kind() == reflect.Int8   }
func (val *Value) IsInt16 () bool { return val.Kind() == reflect.Int16  }
func (val *Value) IsInt32 () bool { return val.Kind() == reflect.Int32  }
func (val *Value) IsInt64 () bool { return val.Kind() == reflect.Int64  }

func (val *Value) Call(args ...any) (res []reflect.Value, err error) { return val.call(Map[any, reflect.Value](reflect.ValueOf)(args), false) }

func (val *Value) call(inputs []reflect.Value, fromPrevious bool) (res []reflect.Value, err error) {
	if val.Kind() != reflect.Func {
		if val.Then == nil || !val.Then.IsFunc() {
			err = fmt.Errorf("attempting to call a non-function [%s] as a function", val.String())
			return
		}
		return val.Then.call([]reflect.Value{val.Value}, false)
	}

	if len(inputs) < 2 {
		res = val.Value.Call(inputs)
		if val.Then != nil { return val.Then.call(res, true) }
		return res, nil
	}

	if fromPrevious {
		switch inp := inputs[len(inputs)-1].Interface().(type) {
		case error: if inp != nil { return nil, inp }; inputs = inputs[:len(inputs)-1]
		case bool: if !inp        { return nil, fmt.Errorf("in call to [%s] the last argument [ok != true]", val.String()) }; inputs = inputs[:len(inputs)-1] }
	}

	res = val.Value.Call(inputs)
	if val.Then != nil { return val.Then.call(res, true) }
	return res, nil
}

func Function(input any) (res *Value) { return &Value{ Value: reflect.ValueOf(input) } }

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
func Filter[T any](fns ...Pred[T]) func([]T) []T {
	fn := PredAnd(fns...)
	return func(args []T) (res []T) {
		for _, arg := range args {
			if fn(arg) {
				res = append(res, arg)
			}
		}
		return res
	}
}

func FilterMap[I, O any](fn func(I) Option[O]) func([]I) []O {
	return func(arr []I) []O {
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

func Over[I, O any, Fn Flatter[I, O] | Mapper[I, O] | TTVar[I, O] | AVar[O] | Say | TVar[I, O] | Shout[I] | Pred[I]](arg Fn) (res func([]I) []O) {
	switch fn := any(arg).(type) {
	case Say     : return Cast[func([]I) []O](Each(func(t I) { fn(t) })).Must()
	case Shout[I]: return Cast[func([]I) []O](Each(fn)).Must()

	case Pred   [I]   : return Cast[func([]I) []O](Filter(fn)).Must()
	case Flatter[I, O]: return FilterMap(fn)

	case TVar  [I, O]: return Map[I, O](func(t I) O { return fn(t) })
	case AVar  [O]   : return Map[I, O](func(t I) O { return fn(t) })
	case TTVar [I, O]: return Map[I, O](func(t I) O { return fn(t) })
	case Mapper[I, O]: return Map[I, O](fn)
	default          : return
	}
}

func Each[T any, Arr []T](fn func(T)) func(Arr) Arr {
	return func(arr Arr) Arr {
		for _, value := range arr {
			fn(value)
		}
		return arr
	}
}


// Map applies the function to each argument and returns the results.
func Map[I, O any, Fn MapFn[I, O]](arg Fn) func([]I) []O {
	var fn func(I) O
	switch fun := any(arg).(type) {
	case TVar[I, O]    : fn = func(t I) O { return fun(t) }
	case AVar[O]       : fn = func(t I) O { return fun(t) }
	case TTVar[I, O]   : fn = func(t I) O { return fun(t) }
	case Mapper[I, O]  : fn = fun }

	return func(args []I) (res []O) {
		res = make([]O, 0, len(args))
		for _, arg := range args { res = append(res, fn(arg)) }
		return res
	}
}

func FlatMap[I, O any](fn func(I) []O) func([]I) []O {
	return func(args []I) (res []O) {
		for _, arg := range args {
			res = append(res, fn(arg)...)
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

// FindFirst returns the first value which passes the [Predicate].
func FindFirst[T any](fns ...Pred[T]) func([]T) Option[T] {
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

func ABC(values ...any) (res *Value, ok bool){
	var top *Value
	var last *Value
	for _, val := range values {
		fc := Function(val)
		if top == nil { top = fc; last = top; continue }
		last.Then = fc
		last = fc
	}
	return top, true
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

func Cat[I, T, O any](a func(I) T, b func(T) O) func(I) O { return func(c I) O { return b(a(c)) } }

func String[T byte | []byte | rune | []rune | string](arg T) string { return string(arg) }

func IfCat[I, T, O any](actor func(I) (T, bool), transformer func(T) O) func(I) (O, bool) {
	return func(i I) (res O, ok bool) {
		var b T
		b, ok = actor(i)
		if ok { res = transformer(b) }
		return
	}
}

func IfCat2[I1, I2, T, O any](actor func(I1, I2) (T, bool), transformer func(T) O) func(I1, I2) (O, bool) {
	return func(i1 I1, i2 I2) (res O, ok bool) {
		var b T
		b, ok = actor(i1, i2)
		if ok { res = transformer(b) }
		return
	}
}

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

func Includes[K comparable](inclusive bool) func(args ...K) func([]K) bool {
	return func(args ...K) func([]K) bool {
		var pred Pred[K]
		// if inclusive { pred = Is(args...) } else { pred = IsEvery(args...) }
		pred = Is(args...)
		return func(arr []K) bool { return slices.ContainsFunc(arr, pred) }
	}
}

func Vals[K comparable, V any](m map[K]V) (res []V) {
	if m == nil { return }
	for _, v := range m {res = append(res, v)}
	return
}

func Keys[K comparable, V any](m map[K]V) (res []K) {
	if m == nil { return }
	for k := range m {res = append(res, k)}
	return
}
