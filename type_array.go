package blume

type Array[A any] struct{ Val []A }

func (a Array[A]) Map(fns ...FnA[A, A]) Array[A]       { return Array[A]{Map(Pipe(fns...))(a.Val)} }
func (a Array[A]) First(fns ...FnA[A, bool]) Option[A] { return First(fns...)(a.Val) }
func (a Array[A]) Reduce(fn FnB[A, A, A], init A) A     { return Reduce(fn, init)(a.Val) }
func (a Array[A]) Filter(fns ...FnA[A, bool]) Array[A] { return Array[A]{Filter(fns...)(a.Val)} }

func (a Array[A]) Append(b ...A) Array[A]  { return Array[A]{append(a.Val, b...)} }
func (a Array[A]) Prepend(b ...A) Array[A] { return Array[A]{append(b, a.Val...)} }
func (a Array[A]) Len() int                { return len(a.Val) }
func (a Array[A]) Values() []A             { return a.Val }

// Slice !!!UNTESTED!!!
func (a Array[A]) Slice(from, to int) Array[A] {
	if to < 0 { to = len(a.Val) + to }
	if from < 0 { from = len(a.Val) + from }
	if from > to { to, from = from, to }
	to = Clamp(0, len(a.Val))(to)
	from = Clamp(0, len(a.Val))(from)

	return Array[A]{a.Val[from:to]}
}

// Slice !!!UNTESTED!!!
func (a *Array[A]) Shift() Option[A] {
	if len(a.Val) == 0 { return None[A]() }
	res := a.Val[0]
	a.Val = a.Val[1:]
	return Some(res)
}

// Slice !!!UNTESTED!!!
func (a *Array[A]) Pop() Option[A] {
	if len(a.Val) == 0 { return None[A]() }
	res := a.Val[len(a.Val)-1]
	a.Val = a.Val[:len(a.Val)-1]
	return Some(res)
}

func ToArray[A any](a []A) Array[A] { return Array[A]{a} }

func Pair[A any](arr []A) Array[[]A] {
	pairs := [][]A{}
	var i int
	for i = 0; i < len(arr); i += 2 {
		cur := []A{}
		if i+1 <= len(arr) { cur = append(cur, arr[i]) }
		if i+2 <= len(arr) { cur = append(cur, arr[i+1]) }
		pairs = append(pairs, cur)
	}
	return ToArray(pairs)
}
