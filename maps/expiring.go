package maps

import (
	"iter"
	"time"

	. "github.com/periaate/blume"
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
func (em *Expiring[K, V]) Get(k K) Option[V] {
	it := em.Sync.Get(k)
	if !it.Ok {
		return None[V]()
	}
	if isExpired(it.Value.Expires) {
		em.mut.Lock()
		em.lockless_del(k)
		em.mut.Unlock()
		return None[V]()
	}
	return Some(it.Value.Value)
}

// Set adds or updates a value in the map for a given key.
func (em *Expiring[K, V]) Set(k K, v V, dur time.Duration) Option[V] {
	expires := time.Now().Add(dur)
	if isExpired(expires) {
		return None[V]()
	}
	return Some(em.Sync.Set(k, ExpItem[V]{v, expires}).Value)
}

func (em *Expiring[K, V]) Del(k K) bool { return em.Sync.Del(k) }
