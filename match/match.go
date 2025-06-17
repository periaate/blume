package match

import (
	"fmt"
	"reflect"
	"slices"
	"strings"

	"github.com/periaate/blume"
)

type Action int

const (
	Nothing Action = iota
	Replaces
	Deletes
	Stops

	Keep
	Skip
)

type SplitArguments[T any] struct {
	Window[T]
	Size[T]
	Eq[T]
	Action
}

type act[T any] struct {
	I T
	A Action
}

type actFn[T any] struct {
	S int
	F func(T) bool
	A Action
}

func Until[T any](until bool, fn func(T) bool, action Action) (res Match[Window[T], SplitResult[T]]) {
	return func(w Window[T]) (res SplitResult[T]) {
		var i int = 2
		var last_v T
		for {
			val, ok := w(i)
			if !ok {
				res.Size = i
				res.Ok = false
				return
			}
			fmt.Println(i, fn(val))
			if fn(val) == until {
				res.Size = i-1
				res.Ok = true
				res.Action = action
				res.Result = last_v
				return
			}
			last_v = val
			i += 1
		}
	}
}


func Act[T any](v T, a Action) act[T] { return act[T]{v, a} }
func ActFn[T any](s int, fn func(T) bool, a Action) actFn[T] { return actFn[T]{s, fn, a} }

type SplitResult[T any] struct {
	Result T
	Size int
	Ok bool
	Action
}

type Window[T any] func(size int) (result T, ok bool)
type Size  [T any] func(item T) (size int)
type Eq    [T any] func(T, T) bool

type Match [Source, Result any] func(Source) Result

func (s Size[T]) Eq(a, b T) int { return s(b)-s(a) }

func Is[T any](sizer Size[T], eq Eq[T], items ...act[T]) (res Match[Window[T], SplitResult[T]]) {
	slices.SortFunc(items, func(a, b act[T]) int { return sizer.Eq(a.I, b.I) })
	return func(src Window[T]) (res SplitResult[T]) {
		for _, item := range items {
			res.Size = sizer(item.I)
			res.Result, res.Ok = src(res.Size)
			if !res.Ok { continue }
			if !eq(item.I, res.Result) { continue }
			res.Action = item.A
			return res
		}
		return SplitResult[T]{}
	}
}

func IsBy[T any](items ...actFn[T]) (res Match[Window[T], SplitResult[T]]) {
	slices.SortFunc(items, func(a, b actFn[T]) int { return b.S-a.S })
	return func(src Window[T]) (res SplitResult[T]) {
		for _, item := range items {
			res.Size = item.S
			res.Result, res.Ok = src(res.Size)
			if !res.Ok { continue }
			if !item.F(res.Result) { continue }
			res.Action = item.A
			return res
		}
		return SplitResult[T]{}
	}
}

func prettyPrintOfType[T any](t T) string {
	switch val := any(t).(type) {
	case string: return `"`+blume.Replace(`\`, `\\`, `"`, `\"`)(val)+`"`
	default: 
		typ := reflect.TypeOf(t)
		if typ.Kind() == reflect.Slice || typ.Kind() == reflect.Array {
			return prettyPrintSlice[any](any(t).([]any))
		}
		if typ.Comparable() { return fmt.Sprint(t) }
		return fmt.Sprintf("<%s>{%v}", typ.Name(), t)
	}
}

func prettyPrintSlice[T any](arr []T) string {
	return `[`+strings.Join(blume.Over[T, string](prettyPrintOfType[T])(arr), ", ")+`]`
}

func Split[Arr, Item any](itr Iter[Arr, Item], match Match[Window[Arr], SplitResult[Arr]]) (result []Arr) {
	startOfSegment := 0

	for {
		i := itr.I()
		res := match(func(size int) (Arr, bool) { return itr.Slice(i, i+size) })
		if !res.Ok {
			if _, ok := itr.Step(1); !ok {
				break
			}
			continue
		}

		if i > startOfSegment {
			preMatch, ok := itr.Slice(startOfSegment, i)

			if ok && !IsZero(preMatch) { result = append(result, preMatch) }
		}

		if res.Action == Keep && !IsZero(res.Result) { result = append(result, res.Result) }

		itr.Step(res.Size)
		startOfSegment = itr.I()
	}

	lastSegment, ok := itr.Slice(startOfSegment, itr.I())
	if ok {
		// if !IsZero(lastSegment) {
		// v := reflect.ValueOf(lastSegment)
		// if v.IsValid() && v.Len() > 0 {
		if !IsZero(lastSegment) { result = append(result, lastSegment) }
	}

	return result
}

func IsZero[T any](val T) (ok bool) {
	var def T
	return reflect.DeepEqual(def, val)
}
