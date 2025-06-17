package match

import (
	"fmt"
	"iter"

	"github.com/periaate/blume"
)

type Iter[Arr, Item any] interface {
	Index(idx int) (Item, bool)
	Slice(from, to int) (Arr, bool)
	Step(n int) (Item, bool)
	Window(n int) (Arr, bool)
	Iter() iter.Seq2[int, Item]
	Reverse() Iter[Arr, Item]
	I() int
}

func (itr anyIterator[Src, Arr, Item]) I() int { return itr.idx }
func (itr sliceIterator[T]) I() int { return itr.idx }

func newIter[T any](input []T) Iter[[]T, T] { return &sliceIterator[T]{ arr: input }}

type sliceIterator[T any] struct {
	arr []T
	idx int
	rev bool
}

type anyIterator[Src, Arr, Item any] struct {
	src Src
	ranges Range[Src, Arr]
	indexes Index[Src, Item]
	idx int
	rev bool
}

func (itr *sliceIterator[T])            Reverse() Iter[[]T, T]    { itr.rev=!itr.rev; return itr }
func (itr *anyIterator[Src, Arr, Item]) Reverse() Iter[Arr, Item] { itr.rev=!itr.rev; return itr }

func (itr *anyIterator[Src, Arr, Item]) Slice(from, to int) (Arr, bool) { return itr.ranges(itr.src, from, to) }
func (itr *anyIterator[Src, Arr, Item]) Index(idx int) (Item, bool) { return itr.indexes(itr.src, idx) }
func (itr *anyIterator[Src, Arr, Item]) Window(n int) (Arr, bool) { return itr.ranges(itr.src, itr.idx, itr.idx+n) }

var _ Iter[[]any, any] = &sliceIterator[any]{}
var _ Iter[[]any, any] = &anyIterator[any, []any, any]{}

func (itr *sliceIterator[T]) Index(n int) (res T, ok bool) { if len(itr.arr) > n { return itr.arr[n], true }; return }
func (itr *sliceIterator[T]) Window(n int) (res []T, ok bool) { return itr.Slice(itr.idx, itr.idx+n) }

func (itr *sliceIterator[T]) Slice(from, to int) (res []T, ok bool) {
	s, l := min(from, to), max(from, to)
	if s < 0 || l < s || len(itr.arr) < l { return }
	return itr.arr[s:l], true
}

func (itr *sliceIterator[T]) Step(n int) (res T, ok bool) {
	if n == 0 { n=1 }
	res, ok = itr.Index(itr.idx)
	if ok { switch itr.rev {
	case false:  itr.idx+=n
	case true: itr.idx-=n } }
	return
}

func (itr *anyIterator[Src, Arr, Item]) Step(n int) (res Item, ok bool) {
	if n == 0 { n=1 }
	res, ok = itr.Index(itr.idx)
	if ok { switch itr.rev {
	case false:  itr.idx+=n
	case true: itr.idx-=n } }
	return
}

type Range[Src, Arr any] func(src Src, from int, to int) (result Arr, ok bool)
type Index[Src, Item any] func(src Src, idx int) (result Item, ok bool)

func ToIter[Arr string | []Item, Item any | rune | byte | string](input Arr) (res Iter[Arr, Item], err error) {
	var item Item
	switch value := any(input).(type) {
	case string:
		switch any(item).(type) {
		case rune: return ToIterBy(
			[]rune(value),
			blume.Get[rune],
			func(src []rune, from, to int) (result string, ok bool) {
				val, ok := blume.Slice(src, from, to)				
				if !ok { return }
				return string(val), ok
			},
		).(Iter[Arr, Item]), nil
		case string: return ToIterBy(
			[]rune(value),
			blume.IfCat2(blume.Get[rune], blume.String),
			func(src []rune, from, to int) (result string, ok bool) {
				val, ok := blume.Slice(src, from, to)				
				if !ok { return }
				return string(val), ok
			},
		).(Iter[Arr, Item]), nil
		case byte: return ToIterBy(
			[]byte(value),
			blume.Get[byte],
			func(src []byte, from, to int) (result string, ok bool) {
				val, ok := blume.Slice(src, from, to)				
				if !ok { return }
				return string(val), ok
			},
		).(Iter[Arr, Item]), nil
		default: return res, fmt.Errorf("illegal invariant of string: Element type must be either rune, string, or byte")
		}
	case []Item: return newIter(value).(Iter[Arr, Item]), nil
	default: return res, fmt.Errorf("impossible or illegal invariant; is neither string nor slice type")
	}
}

func ToIterBy[Src, Arr, Item any](input Src, idx Index[Src, Item], rng Range[Src, Arr]) Iter[Arr, Item] {
	return &anyIterator[Src, Arr, Item]{
		src: input,
		indexes: idx,
		ranges: rng,
	}
}


func Shift[T any](arr []T) (res T, out []T, ok bool) { if len(arr) > 0 { res, out, ok = arr[0],          arr[1:],        true }; return }
func Pop[T any]  (arr []T) (res T, out []T, ok bool) { if len(arr) > 0 { res, out, ok = arr[len(arr)-1], arr[:len(arr)-1], true }; return }

func (itr *sliceIterator[T]) Iter() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for {
			el, ok := itr.Step(0)
			if !ok { return }
			if !yield(itr.idx-1, el) { return }
		}
	}
}

func (itr *anyIterator[Src, Arr, Item]) Iter() iter.Seq2[int, Item] {
	return func(yield func(int, Item) bool) {
		for {
			i := itr.idx
			el, ok := itr.Step(0)
			if !ok { return }
			if !yield(i, el) { return }
		}
	}
}
