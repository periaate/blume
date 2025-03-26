package blume

func Through[A any](fn func(A)) func(A) A { return func(arg A) A { fn(arg); return arg } }
func T[A any](ok bool, a A, b A) A {
	if ok {
		return a
	} else {
		return b
	}
}
func Thunk[A, B any](val A) func(_ B) A { return func(_ B) A { return val } }
