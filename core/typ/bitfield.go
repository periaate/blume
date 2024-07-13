package typ

import "blume/core/num"

func AND[N num.UInteger](a, b N) N    { return a & b }
func NOT[N num.UInteger](a, b N) N    { return a &^ b }
func OR[N num.UInteger](a, b N) N     { return a | b }
func XOR[N num.UInteger](a, b N) N    { return a ^ b }
func HAS[N num.UInteger](a, b N) bool { return (a & b) != 0 }

func Include[N num.UInteger](src N, args ...N) N {
	for _, v := range args {
		src |= v
	}
	return src
}

type BitField[K comparable, N num.UInteger] struct {
	Aliases map[K]N
}

func (bf *BitField[K, N]) WithKey(args ...K) (res N) {
	for _, key := range args {
		if v, ok := bf.Aliases[key]; ok {
			res = OR(res, v)
		}
	}
	return
}

func (bf *BitField[K, N]) Alias(n N, keys ...K) {
	for _, key := range keys {
		bf.Aliases[key] |= n
	}
}
