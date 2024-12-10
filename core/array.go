package core

type Arr[A any] struct{ val []A }

var _ Array[any] = Arr[any]{}

func (a Arr[A]) Append(b ...A) Arr[A]                { return Arr[A]{append(a.val, b...)} }
func (a Arr[A]) Map(fns ...Monadic[A, A]) Array[A]   { return Arr[A]{Map(Pipe(fns...))(a.val)} }
func (a Arr[A]) Filter(fns ...Predicate[A]) Array[A] { return Arr[A]{Filter(fns...)(a.val)} }
func (a Arr[A]) First(fns ...Predicate[A]) Option[A] { return First(fns...)(a.val) }
func (a Arr[A]) Reduce(fn Dyadic[A, A, A], init A) A { return Reduce(fn, init)(a.val) }
func (a Arr[A]) Len() int                            { return len(a.val) }
func (a Arr[A]) Values() []A                         { return a.val }

func ToArray[A any](a []A) Array[A]                     { return Arr[A]{a} }
func ArrayFrom[A, B any](a Arr[A], fn func(A) B) Arr[B] { return Arr[B]{Map(fn)(a.Values())} }
