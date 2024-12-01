package maps

import (
	"iter"
	"time"
)

type item[V any] struct {
	v V
	t time.Time
}

// Expiring is a thread safe map where values have expiration dates.
// Expiring does not automatically clear expired items, rather, they are deleted on Get.
type Expiring[K comparable, V any] struct {
	*Sync[K, item[V]]
}

// Keys returns a sequence of keys in the map.
func (em *Expiring[K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		em.mut.RLock()
		defer em.mut.RUnlock()
		for k := range em.val {
			if !yield(k) {
				return
			}
		}
	}
}

// Vals returns a sequence of values in the map.
func (em *Expiring[K, V]) Vals() iter.Seq[V] {
	return func(yield func(V) bool) {
		em.mut.RLock()
		defer em.mut.RUnlock()
		for _, v := range em.val {
			if !yield(v.v) {
				return
			}
		}
	}
}

// Iter returns a sequence of key-value pairs in the map.
func (em *Expiring[K, V]) Iter() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		em.mut.RLock()
		defer em.mut.RUnlock()
		for k, v := range em.val {
			if !yield(k, v.v) {
				return
			}
		}
	}
}

// NewSync initializes and returns a new Sync.
func NewExpiring[K comparable, V any]() *Expiring[K, V] {
	return &Expiring[K, V]{NewSync[K, item[V]]()}
}

func isExpired(t time.Time) bool { return t.Before(time.Now()) }

// Get retrieves a value by key. It returns the value and a boolean indicating if the key exists.
func (em *Expiring[K, V]) Get(k K) (res V, ok bool) {
	it, ok := em.Sync.Get(k)
	if !ok {
		return
	}

	if isExpired(it.t) {
		ok = em.Del(k)
		return
	}

	res = it.v
	return
}

// Set adds or updates a value in the map for a given key.
func (em *Expiring[K, V]) Set(k K, v V, expiration time.Time) (ok bool) {
	if isExpired(expiration) {
		return
	}

	return em.Sync.Set(k, item[V]{v, expiration})
}

// Del removes a value by key. It returns a boolean indicating if the key was found and removed.
func (em *Expiring[K, V]) Del(k K) (ok bool) {
	return em.Sync.Del(k)
}
