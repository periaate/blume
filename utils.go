package blume

import (
	"fmt"
	"iter"
	"os"
	"reflect"
)

type Iterator[A, T any] interface {
	Index(idx int) (T, bool)
	Slice(from, to int) (A, bool)
	Step(n int) (T, bool)
	Window(n int) (A, bool)
	Iter() iter.Seq2[int, T]
	Len() int
}

func newStringRuneIter(input string) Iterator[string, rune] {
	return stringRuneIterator{sliceIterator[rune]{arr: []rune(input)}}
}

func newStringByteIter(input string) Iterator[string, byte] {
	return stringByteIterator{sliceIterator[byte]{arr: []byte(input)}}
}

func newIter[T any](input []T) Iterator[[]T, T] { return sliceIterator[T]{ arr: input }}

type stringRuneIterator struct { sliceIterator[rune] }
type stringByteIterator struct { sliceIterator[byte] }

func (itr stringRuneIterator) Slice(from, to int) (res string, ok bool) { val, ok := itr.sliceIterator.Slice(from, to); return string(val), ok }
func (itr stringByteIterator) Slice(from, to int) (res string, ok bool) { val, ok := itr.sliceIterator.Slice(from, to); return string(val), ok }
func (itr stringRuneIterator) Window(n int) (res string, ok bool) { return itr.Slice(itr.idx, itr.idx+n) }
func (itr stringByteIterator) Window(n int) (res string, ok bool) { return itr.Slice(itr.idx, itr.idx+n) }

type sliceIterator[T any] struct {
	arr []T
	idx int
}

var _ Iterator[[]any, any] = sliceIterator[any]{}

func (itr sliceIterator[T]) Len() int { return len(itr.arr) }

func (itr sliceIterator[T]) Index(n int) (res T, ok bool) { if len(itr.arr) > n { return itr.arr[n], true }; return }
func (itr sliceIterator[T]) Window(n int) (res []T, ok bool) { return itr.Slice(itr.idx, itr.idx+n) }
func (itr sliceIterator[T]) Slice(from, to int) (res []T, ok bool) {
	s, l := min(from, to), max(from, to)
	if s < 0 || l < s || len(itr.arr) < l { return }
	return itr.arr[s:l], true
}
func (itr sliceIterator[T]) Step(n int) (T, bool) { itr.idx+=n; return itr.Index(itr.idx) }


func Index[T any](arr []T, i int) (res Option[T]) {
	if i < 0         { i = len(arr) + i }
	if i < 0         { return res.Fail() }
	if i >= len(arr) { return res.Fail() }
	return res.Pass(arr[i])
}

func Iter[El any | rune | byte | string, Arr string | []El](input Arr) (res Iterator[El, Arr], err error) {
	var out El
	switch value := any(input).(type) {
	case string:
		switch any(out).(type) {
		case rune: return newStringRuneIter(value).(Iterator[El, Arr]), nil
		case byte: return newStringByteIter(value).(Iterator[El, Arr]), nil
		default: return res, fmt.Errorf("illegal invariant of string: Element type must be either rune, string, or byte")
		}
	case []El: return newIter(value).(Iterator[El, Arr]), nil
	default: return res, fmt.Errorf("impossible or illegal invariant; is neither string nor slice type")
	}
}


func Shift[T any](arr []T) (res T, out []T, ok bool) { if len(arr) > 0 { res, out, ok = arr[0],          arr[1:],        true }; return }
func Pop[T any]  (arr []T) (res T, out []T, ok bool) { if len(arr) > 0 { res, out, ok = arr[len(arr)-1], arr[:len(arr)], true }; return }

func (e Either[A, B]) Is(value A) bool {
	if !e.IsOk()                           { return false }
	if !reflect.TypeOf(value).Comparable() { return false }
	return reflect.ValueOf(e.Value).Equal(reflect.ValueOf(value))
}

type Op[I, O any] interface { func(I) func(I) O | func(I, I) O }

func Pattern[T any, Fn Op[T, bool]](pred Fn, args ...T) (ok bool) {
	if len(args) == 0 { return }
	var op func(T, T) bool
	switch fn := any(pred).(type) {
	case func(T) Pred[T]: op = func(t1, t2 T) bool { return fn(t2)(t1) }
	case func(T, T) bool: op = fn }
	for i, arg := range args[1:] {
		if !op(args[i], arg) { return }
	}
	return true
}

func (itr sliceIterator[T]) Iter() iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i := range len(itr.arr)-itr.idx {
			el, ok := itr.Index(i)
			if !ok { return }
			if !yield(i+itr.idx, el) { break }
		}
	}
}

// Exit the program with a console log
func Exit(args ...any) {
	fmt.Printf("%s\n", Join(" ")(args))
	os.Exit(1)
}

func ExitWith(n int, args ...any) { fmt.Printf("%s", Join(" ")(args)); os.Exit(n) }

func ExitsWith[A any](n int) func(arg A) A { return func(arg A) A { ExitWith(n, arg); return arg } }

func OrExit[A, B any](either Either[A, B], args ...any) (res A) {
	if !either.IsOk() { Exit(fmt.Sprintf("%s [%v]", fmt.Sprint(args...), either.Other)) }
	return either.Value
}

func OrExits[A, B any](either Either[A, B]) (res A) {
	if !either.IsOk() { Exit(fmt.Sprintf("%v", either.Other)) }
	return either.Value
}
