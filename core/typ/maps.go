package typ

func Invert[K, V comparable](m map[K]V) (res map[V][]K) {
	if m == nil {
		return
	}

	res = map[V][]K{}

	for key, val := range m {
		if _, ok := res[val]; !ok {
			res[val] = make([]K, 1)
		}
		res[val] = append(res[val], key)
	}

	return
}

func InvertAny[K, C comparable, V any](m map[K]V, fn func(V) C) (res map[C][]K) {
	if m == nil {
		return
	}

	res = map[C][]K{}

	for k, v := range m {
		key := fn(v)
		if _, ok := res[key]; !ok {
			res[key] = make([]K, 1)
		}
		res[key] = append(res[key], k)
	}

	return
}
