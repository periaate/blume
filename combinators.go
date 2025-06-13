package blume

func Through[A any](fn func(A)) func(A) A { return func(arg A) A { fn(arg); return arg } }
func Ignore[A any](fn func(A) A) func(A) { return func(arg A) { fn(arg) } }
func T[A any](ok bool, a A, b A) A {
	if ok {
		return a
	} else {
		return b
	}
}
func Thunk[A, B any](val A) func(_ B) A { return func(_ B) A { return val } }

func V2M[A, B any](fn func(...A) B) func(A) B { return func(arg A) B { return fn(arg) }}

// type Ifs[A any] func(bool) A

// func If[A any]() Ifs[A] { return func(b bool) (res A) { return } }

// func (i Ifs[A]) Then(arg A) Ifs[A] {
// 	return func(input bool) (output A) {
// 		if input { return arg }
// 		return i(input)
// 	}
// }
//
// func (i Ifs[A]) Else(arg A) Ifs[A] {
// 	return func(input bool) (output A) {
// 		if !input { return arg }
// 		return i(input)
// 	}
// }

func If[A any, B ~bool](ok B, a A, b A) A {
	if ok {
		return a
	} else {
		return b
	}
}
