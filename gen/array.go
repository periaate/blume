package gen

import (
	"iter"

	"github.com/periaate/blume/gen/T"
)

func ToArray[A any](a []A) Array[A] { return Array[A]{a} }

type Array[A any] struct {
	val []A
}

func (a Array[A]) Append(b ...A) Array[A] {
	return Array[A]{append(a.val, b...)}
}

func (a Array[A]) GetPop() T.Result[A] {
	var zero A
	if len(a.val) == 0 {
		return T.Results(zero, "empty array")
	}
	return T.Results(a.val[len(a.val)-1], nil)
}

func (a Array[A]) GetShift() T.Result[A] {
	var zero A
	if len(a.val) == 0 {
		return T.Results(zero, "empty array")
	}
	return T.Results(a.val[0], nil)
}

func (a Array[A]) Map(fns ...T.Monadic[A, A]) Array[A] {
	return Array[A]{Map(Pipe(fns...))(a.val)}
}

func (a Array[A]) Filter(fns ...T.Predicate[A]) Array[A] {
	return Array[A]{Filter(fns...)(a.val)}
}

func (a Array[A]) First(fns ...T.Predicate[A]) T.Result[A] {
	return First(fns...)(a.val)
}

func (a Array[A]) Reduce(fn T.Dyadic[A, A, A], init A) A {
	return Reduce(fn, init)(a.val)
}

func (a Array[A]) Values() iter.Seq[A] {
	return func(yield func(A) bool) {
		for _, v := range a.val {
			if !yield(v) {
				return
			}
		}
	}
}

func (a Array[A]) Iter() iter.Seq2[int, A] {
	return func(yield func(int, A) bool) {
		for i, v := range a.val {
			if !yield(i, v) {
				return
			}
		}
	}
}

func (a Array[A]) Len() int   { return len(a.val) }
func (a Array[A]) Array() []A { return a.val }

func ArrayFrom[A, B any](a Array[A], fn func(A) B) Array[B] { return Array[B]{Map(fn)(a.Array())} }
