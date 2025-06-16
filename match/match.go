package match

import (
	"slices"
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

func Act[T any](v T, a Action) act[T] { return act[T]{v, a} }


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

func (s Size[T]) Eq(a, b T) int { return s(a)-s(b) }

func Is[T any](sizer Size[T], eq Eq[T], items ...act[T]) (res Match[Window[T], SplitResult[T]]) {
	slices.SortFunc(items, func(a, b act[T]) int { return sizer.Eq(a.I, b.I) })
	return func(src Window[T]) (res SplitResult[T]) {
		for _, item := range items {
			res.Size = sizer(item.I)
			res.Result, res.Ok = src(res.Size)
			if !res.Ok { continue }
			res.Action = item.A
			return res
		}
		return SplitResult[T]{}
	}
}

// func Rule[T any](match Match[T], action Action) Pattern[T] {
// }

func Split[Arr, Item any](itr Iter[Arr, Item], match Match[Window[Arr], SplitResult[Arr]]) (result []Arr) {
	lastI := 0
	for i, _ := range itr.Iter() {
		res := match(itr.Window)
		if !res.Ok { continue }
		val, ok := itr.Window(-lastI)
		if ok { result = append(result, val) }
		switch res.Action {
		case Keep: result = append(result, res.Result)
		case Skip:
		}
		itr.Step(res.Size)
		lastI = i+res.Size
	}
	return
}


