package blume

type Array[A any] struct{ Value []A }

type Length int

func (l Length) Is(i int) bool { return int(l) == i }
func (l Length) Gt(i int) bool { return int(l) > i }
func (l Length) Lt(i int) bool { return int(l) < i }
func (l Length) Ge(i int) bool { return int(l) >= i }
func (l Length) Le(i int) bool { return int(l) <= i }
func (l Length) Eq(i int) bool { return int(l) == i }
func (l Length) Ne(i int) bool { return int(l) != i }

func (arr Array[A]) Len() Length { return Length(len(arr.Value)) }

func (arr Array[A]) Get(i int) (res Option[A]) {
	if i < 0 {
		i = len(arr.Value) + i
	}
	if i < 0 {
		return res.Fail()
	}
	if i >= len(arr.Value) {
		return res.Fail()
	}
	return res.Pass(arr.Value[i])
}

func Arr[A any](args ...A) Array[A] { return Array[A]{Value: args} }
func ToArray[A any](a []A) Array[A] { return Array[A]{a} }

func (arr Array[A]) Filter(fn Pred[A]) Array[A] {
	res := []A{}
	for _, val := range arr.Value {
		if fn(val) {
			res = append(res, val)
		}
	}
	return Array[A]{Value: res}
}

func (arr Array[A]) Filter_map(fn func(A) Option[A]) Array[A] {
	res := []A{}
	for _, val := range arr.Value {
		if ret := fn(val); ret.Other {
			res = append(res, ret.Value)
		}
	}
	return Array[A]{Value: res}
}

func (arr Array[A]) First(fn Pred[A]) (res Option[A]) {
	for _, val := range arr.Value {
		if fn(val) {
			return res.Pass(val)
		}
	}
	return res.Fail()
}

func (arr Array[A]) Map(fn func(A) A) Array[A] {
	res := make([]A, len(arr.Value))
	for i, val := range arr.Value {
		res[i] = fn(val)
	}
	return Array[A]{Value: res}
}

func (arr Array[A]) Each(fn func(A)) Array[A] {
	for _, value := range arr.Value {
		fn(value)
	}
	return arr
}
