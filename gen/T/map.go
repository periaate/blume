package T

type Map[K comparable, V any] struct {
	M map[K]V
}

func (m Map[K, V]) Contains(k K) bool {
	_, ok := m.M[k]
	return ok
}

func (m Map[K, V]) Get(k K) (V, bool) {
	v, ok := m.M[k]
	return v, ok
}

func (m Map[K, V]) Set(k K, v V) { m.M[k] = v }
func (m Map[K, V]) Del(k K)      { delete(m.M, k) }
