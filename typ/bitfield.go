package typ

import "github.com/periaate/blume/gen"

// AND bitwise operation.
func AND[N Unsigned](a, b N) N { return a & b }

// NOT bitwise operation.
func NOT[N Unsigned](a, b N) N { return a &^ b }

// OR bitwise operation.
func OR[N Unsigned](a, b N) N { return a | b }

// XOR bitwise operation.
func XOR[N Unsigned](a, b N) N { return a ^ b }

// HAS checks if a bit is set.
func HAS[N Unsigned](a, b N) bool { return (a & b) != 0 }

// // Include ORs together variadic bitfields.
// func Include[N Unsigned](src N, args ...N) N {
// 	for _, v := range args {
// 		src |= v
// 	}
// 	return src
// }

// type BitField[K comparable, N Unsigned] struct {
// 	Aliases map[K]N
// }
//
// func (bf *BitField[K, N]) WithKey(args ...K) (res N) {
// 	for _, key := range args {
// 		if v, ok := bf.Aliases[key]; ok {
// 			res = OR(res, v)
// 		}
// 	}
// 	return
// }
//
// func (bf *BitField[K, N]) Alias(n N, keys ...K) {
// 	for _, key := range keys {
// 		bf.Aliases[key] |= n
// 	}
// }

func inc[N Unsigned](a ...N) gen.Monadic[N, N] {
	return func(b N) (res N) {
		for _, v := range a {
			res = OR(b, v)
		}
		return
	}
}
