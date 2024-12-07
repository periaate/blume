package lazy

import "github.com/periaate/blume/gen/T"

func Niladic[A any](fn T.Niladic[A]) T.Niladic[A] {
	var isLoaded bool
	var value A
	return func() A {
		if isLoaded {
			return value
		}
		value = fn()
		isLoaded = true
		return value
	}
}

func Monadic[A comparable, B any](fn T.Monadic[A, B]) T.Monadic[A, B] {
	cache := map[A]B{}
	return func(input A) (res B) {
		res, ok := cache[input]
		if !ok {
			res = fn(input)
			cache[input] = res
		}
		return res
	}
}
