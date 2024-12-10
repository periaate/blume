package gen

// Invert creates and inverted index of the given map.
func Invert[K, V comparable](m map[K]V) (res map[V][]K) {
	if m == nil { return }

	res = map[V][]K{}

	for key, val := range m {
		if _, ok := res[val]; !ok { res[val] = make([]K, 1) }
		res[val] = append(res[val], key)
	}

	return
}

// InvertAny creates and inverted index of the given map using a custom function.
func InvertAny[K, C comparable, V any](fn func(V) C) func(map[K]V) map[C][]K {
	return func(m map[K]V) (res map[C][]K) {
		if m == nil { return }

		res = map[C][]K{}

		for k, v := range m {
			key := fn(v)
			if _, ok := res[key]; !ok { res[key] = make([]K, 1) }
			res[key] = append(res[key], k)
		}

		return
	}
}

// Keys returns the keys of the given map.
func Keys[K comparable, V any](m map[K]V) []K {
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Vals returns the values of the given map.
func Vals[K comparable, V any](m map[K]V) []V {
	vals := make([]V, 0, len(m))
	for _, v := range m {
		vals = append(vals, v)
	}
	return vals
}
