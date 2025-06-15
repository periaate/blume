package blume

import (
	"fmt"
	"go/format"
	"iter"
	"os"
)

/*
type Indexable[Array, El] := {
	Array: string -> El: rune | byte
	Array: []El   -> El: any
}
*/

type Iterable[El any | rune | byte, Arr string | []El] struct {
	arr Arr
	indexes func(arr Arr, idx int) (El, bool)
	ranges  func(arr Arr, src, idx int) (Arr, bool)
	idx int
	len int
}

func Index[T any](i int) func(arr []T) (res Option[T]) {
	return func(arr []T) (res Option[T]) {
		if i < 0         { i = len(arr) + i }
		if i < 0         { return res.Fail() }
		if i >= len(arr) { return res.Fail() }
		return res.Pass(arr[i])
	}
}

func Shift[T any](arr []T) (res T, out []T, ok bool) { if len(arr) > 0 { res, out, ok = arr[0],          arr[1:],        true }; return }
func Pop[T any]  (arr []T) (res T, out []T, ok bool) { if len(arr) > 0 { res, out, ok = arr[len(arr)-1], arr[:len(arr)], true }; return }

func (r Either[A, B]) From(args ...any) Either[A, B] { return Pipe[Either[A, B]](args...) }
func (e Either[A, B]) Is(args ...any) bool { return false }

func Get[El any | rune | byte, Arr string | []El](i int, arr Arr) (res Option[El]) {
	if i < 0         { i = len(arr) + i }
	if i < 0         { return res.Fail() }
	if i >= len(arr) { return res.Fail() }

	switch any(arr).(type) {
	case string:
		switch any(out).(type) {
		case rune, string:
		case byte:
			return Pipe[Option[El]](
				Cast[[]byte](arr),
				Index[byte](i),
				Cast[El],
			)
		default: return
		}
	default:
	}

	return res, true
}

// func Window[El any | rune | byte | string, Arr string | []El](index int, size int) func(input Arr) (res Arr, ok bool) {
// 	return func(input Arr) (res Arr, ok bool) {
// 		l := len(arr)
// 		if l == 0 { return }
// 		c := Clamp(0, len(arr))
// 		if start < 0 { start = l+start }
// 		if len(ends) == 0 { return Array[T](arr[c(start):]) }
// 		end := ends[0]
// 		if end < 0 { end = l+end }
// 		return arr[c(start):c(end)]
// 	}
// }

// type Op[T any, O any] func(T, T) O
// func Le[Fn Op[T, O], T, O any](args ...any) Op[int, bool].Le

// Pred[int] { return func(n int) bool { return n <= arg } }

type Op[I, O any] interface {
	func(I) func(I) O | func(I, I) O
}

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

func RangeV(len, src, tar int) (smaller, larger int, ok bool) {
	smaller, larger = min(src, tar), max(src, tar)
	ok = Pattern(Le[int], 0, smaller, larger, len)
	return 
}

func (itr Iterable[El, Arr]) Iter() iter.Seq2[int, El] {
	return func(yield func(int, El) bool) {
		for i := range itr.len-itr.idx {
			el, ok := itr.indexes(itr.arr, i)
			if !ok { return }
			if !yield(i+itr.idx, el) { break }
		}
	}
}

func Iter[El any | rune | byte | string, Arr string | []El](input Arr) (res Iterable[El, Arr], ok bool) {
	res = Iterable[El, Arr]{
		arr: input,
		len: len(input),
	}

	var out El
	switch any(input).(type) {
	case string:
		switch any(out).(type) {
		case rune, string:
			res.indexes = func(input Arr, idx int) (el El, ok bool) {
				inp := any(input).(string)
				if idx < len(inp) { el = any(inp[idx]).(El) }
				return
			}
			res.ranges = func(input Arr, src, size int) (ar Arr, ok bool) {
				inp := any(input).(string)
				s, l, ok := RangeV(len(inp), src, size)
				if !ok { return }
				ar, ok = any(inp[s:l]).(Arr)
				return
			}
		case byte:
			var bar []byte
			if bar, ok = any(input).([]byte); !ok { return }
			res.indexes = func(input Arr, idx int) (el El, ok bool) {
				inp := any(input).(string)
				if idx < len(inp) { el = any(bar[idx]).(El) }
				return
			}
			res.ranges = func(input Arr, src, size int) (ar Arr, ok bool) {
				inp := any(input).([]byte)
				s, l, ok := RangeV(len(inp), src, size)
				if !ok { return }
				ar, ok = any(bar[s:l]).(Arr) 
				return
			}
			res.len = len(bar)
		default: return
		}
	default:
		var arr []El
		if arr, ok = any(input).([]El); !ok { return }
		res.indexes = func(input Arr, idx int) (el El, ok bool) {
			if idx < len(input) { el = arr[idx] }
			return
		}
		res.ranges = func(input Arr, src, size int) (ar Arr, ok bool) {
			s, l, ok := RangeV(len(input), src, size)
			if !ok { return }
			ar, ok = any(arr[s:l]).(Arr) 
			return
		}
	}

	return res, true
}

func (itr *Iterable[El, Arr]) Window(n int) (res Option[Arr]) { return res.Auto(itr.ranges(itr.arr, itr.idx+n, n)) }
func (itr *Iterable[El, Arr]) Peek(n int) (res Option[El]) { return res.Auto(itr.indexes(itr.arr, itr.idx+n)) }
func (itr *Iterable[El, Arr]) Next() (res Option[El]) {
	itr.idx+=1
	return res.Auto(itr.indexes(itr.arr, itr.idx)) 
}

func (itr *Iterable[El, Arr]) Step(n int) (res Option[El]) {
	if Pattern(Le[int], 0, itr.idx+n, itr.len) { return res.Fail() }
	itr.idx+=n
	return res.Auto(itr.indexes(itr.arr, itr.idx)) 
}

func CanIndex[El any | rune, Arr string | ~[]El](idx int, input Arr) (res Option[bool]) {
	var out El
	switch inp := any(input).(type) {
	case string:
		switch any(out).(type) {
		case rune:
			if idx+1 >= len(inp) { return res.Fail() }
			out, ok := any(inp[idx+1]).(El)
			if !ok { return res.Fail() }
			return res.Pass(out)
		}
	}

	return res.Fail()
}

func Next[El any | rune, Arr string | ~[]El](idx int, input Arr) (res Option[El]) {
	var out El
	switch inp := any(input).(type) {
	case string:
		switch any(out).(type) {
		case rune:
			if idx+1 >= len(inp) { return res.Fail() }
			out, ok := any(inp[idx+1]).(El)
			if !ok { return res.Fail() }
			return res.Pass(out)
		}
	}
	res
}


// IsFormat checks whether the input string contains any printf directives.
func IsFormat(str string) bool {
	itr, ok := Iter[rune](str)
	if !ok { return false }
	for i, r := range itr.Iter() {
		if r != '%' { return false }
		itr.Peek(1).IsErr('%')
		if i+1 < len(str) && str[i+1] == '%' {
			i++
			continue
		}
		return true
	}
	return false
}

type exit func(...any)

// Exit the program with a console log
var Exit exit = func(args ...any) {
	fmt.Printf("%s\n", Join(" ")(args))
	os.Exit(1)
}


func ExitWith(n int, args ...any) { fmt.Printf("%s", Join(" ")(args)); os.Exit(n) }

func ExitsWith[A any](n int) func(arg A) A { return func(arg A) A { ExitWith(n, arg); return arg } }

func OrExit[A, B any](either Either[A, B], args ...any) (res A) {
	if !either.IsOk() {
		Exit(fmt.Sprintf("%s [%v]", fmt.Sprint(args...), either.Other))
	}
	return either.Value
}

func OrExits[A, B any](either Either[A, B]) (res A) {
	if !either.IsOk() {
		Exit(fmt.Sprintf("%v", either.Other))
	}
	return either.Value
}
