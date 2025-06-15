package blume

import (
	"fmt"
	"iter"
	"os"
	"reflect"
)

type iterator[A, T any] interface {
	indexes(idx int) (T, bool)
	ranges(src, idx int) (A, bool)
	step() (T, bool)
}

type Iterator[El any | rune | byte, Arr string | []El] struct {
	arr Arr
	indexes func(idx int) (El, bool)
	ranges  func(src, idx int) (Arr, bool)
	idx int
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

func (itr Iterator[El, Arr]) Iter() iter.Seq2[int, El] {
	return func(yield func(int, El) bool) {
		for i := range len(itr.arr)-itr.idx {
			el, ok := itr.indexes(i)
			if !ok { return }
			if !yield(i+itr.idx, el) { break }
		}
	}
}

// Iter creates an Iterator for any iterable primitive type pairing.
func Iter[El any | rune | byte | string, Arr string | []El](input Arr) (res Iterator[El, Arr], err error) {
	res = Iterator[El, Arr]{ arr: input }
	var ok bool
	type genIndexes func(idx int) (El, bool)
	type genRanges func(src int, idx int) (Arr, bool)

	var out El
	switch value := any(input).(type) {
	case string:
		switch any(out).(type) {
		case rune:
			input := []rune(value)
			res.indexes, ok = any(func(idx int) (el rune, ok bool) {
				if idx < len(input) { el = input[idx] }
				return
			}).(genIndexes)
			if !ok {
				err = fmt.Errorf("creating an Iterator for string->rune failed when casting Iterator.indexes from concrete type to generic type; this error case is impossible")
				return
			}

			res.ranges, ok = any(func(src, size int) (ar string, ok bool) {
				s, l, ok := RangeV(len(input), src, size)
				if !ok { return }
				return string(input[s:l]), true
			}).(genRanges)
			if !ok {
				err = fmt.Errorf("creating an Iterator for string->rune failed when casting Iterator.ranges from concrete type to generic type; this error case is impossible")
				return
			}
		case string:
			input := []rune(value)
			res.indexes, ok = any(func(idx int) (el string, ok bool) {
				if idx < len(input) { el = string(input[idx]) }
				return
			}).(genIndexes)
			if !ok {
				err = fmt.Errorf("creating an Iterator for string->string failed when casting Iterator.indexes from concrete type to generic type; this error case is impossible")
				return
			}

			res.ranges, ok = any(func(src, size int) (ar string, ok bool) {
				s, l, ok := RangeV(len(input), src, size)
				if !ok { return }
				return string(input[s:l]), true
			}).(genRanges)
			if !ok {
				err = fmt.Errorf("creating an Iterator for string->string failed when casting Iterator.ranges from concrete type to generic type; this error case is impossible")
				return
			}
		case byte:
			res.indexes, ok = any(func(idx int) (el byte, ok bool) {
				if idx < len(input) { el = value[idx] }
				return
			}).(genIndexes)
			if !ok {
				err = fmt.Errorf("creating an Iterator for string->byte failed when casting Iterator.indexes from concrete type to generic type; this error case is impossible")
				return
			}

			res.ranges, ok = any(func(src, size int) (ar string, ok bool) {
				s, l, ok := RangeV(len(value), src, size)
				if !ok { return }
				return string(value[s:l]), true
			}).(genRanges)
			if !ok {
				err = fmt.Errorf("creating an Iterator for string->byte failed when casting Iterator.ranges from concrete type to generic type; this error case is impossible")
				return
			}
		default: return res, fmt.Errorf("illegal invariant of string: Element type must be either rune, string, or byte")
		}
	case []El:
		name := reflect.TypeOf(out).Name()
		res.indexes, ok = any(func(idx int) (el El, ok bool) {
			if idx < len(input) { el = value[idx] }
			return
		}).(genIndexes)
		if !ok {
			err = fmt.Errorf("creating an Iterator for []%s->%s failed when casting Iterator.indexes from concrete type to generic type; this error case is impossible", name)
			return
		}
		res.ranges, ok = any(func(src, size int) (ar []El, ok bool) {
			s, l, ok := RangeV(len(value), src, size)
			if !ok { return }
			return value[s:l], true
		}).(genRanges)
		if !ok {
			err = fmt.Errorf("creating an Iterator for []%s->%s failed when casting Iterator.ranges from concrete type to generic type; this error case is impossible", name)
			return
		}
	default: return res, fmt.Errorf("impossible or illegal invariant; is neither string nor slice type")
	}

	return res, nil
}

func (itr *Iterator[El, Arr]) Window(n int) (res Option[Arr]) { return res.Auto(itr.ranges(itr.idx+n, n)) }
func (itr *Iterator[El, Arr]) Peek(n int) (res Option[El]) { return res.Auto(itr.indexes(itr.idx+n)) }
func (itr *Iterator[El, Arr]) Next() (res Option[El]) {
	itr.idx+=1
	return res.Auto(itr.indexes(itr.idx))
}

func (itr *Iterator[El, Arr]) Step(n int) (res Option[El]) {
	if Pattern(Le[int], 0, itr.idx+n, len(itr.arr)) { return res.Fail() }
	itr.idx+=n
	return res.Auto(itr.indexes(itr.idx)) 
}

// Exit the program with a console log
func Exit(args ...any) {
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
