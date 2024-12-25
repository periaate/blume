package maps

import (
	"iter"
	"time"
)

type ExpItem[V any] struct {
	Value   V
	Expires time.Time
}

// Expiring is a thread safe map where values have expiration dates.
// Expiring does not automatically clear expired items, rather, they are deleted on Get.
type Expiring[K comparable, V any] struct{ *Sync[K, ExpItem[V]] }

// Iter returns a sequence of key-value pairs in the map.
func (em *Expiring[K, V]) Iter() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		expiredKeys := []K{}
		em.mut.RLock()
		for k, v := range em.values {
			if isExpired(v.Expires) {
				expiredKeys = append(expiredKeys, k)
				continue
			}
			if !yield(k, v.Value) {
				em.mut.RUnlock()
				break
			}
		}

		if len(expiredKeys) == 0 {
			return
		}

		em.mut.Lock()
		defer em.mut.Unlock()
		for _, v := range expiredKeys {
			em.lockless_del(v)
		}
	}
}

// NewSync initializes and returns a new Sync.
func NewExpiring[K comparable, V any]() *Expiring[K, V] {
	exp := &Expiring[K, V]{Sync: NewSync[K, ExpItem[V]]()}
	return exp
}

func isExpired(t time.Time) bool { return t.Before(time.Now()) }

// Get retrieves a value by key. It returns the value and a boolean indicating if the key exists.
func (em *Expiring[K, V]) Get(k K) (res V, ok bool) {
	it, ok := em.Sync.Get(k)
	if !ok {
		return
	}
	if isExpired(it.Expires) {
		em.mut.Lock()
		em.lockless_del(k)
		em.mut.Unlock()
		return
	}
	return
}

// Set adds or updates a value in the map for a given key.
func (em *Expiring[K, V]) Set(k K, v V, dur time.Duration) (ok bool) {
	expires := time.Now().Add(dur)
	if isExpired(expires) {
		return
	}
	em.Sync.Set(k, ExpItem[V]{v, expires})
	return true
}

func (em *Expiring[K, V]) Del(k K) bool { return em.Sync.Del(k) }
