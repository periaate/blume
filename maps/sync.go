package maps

import (
	"iter"
	"sync"
)

func (sm *Sync[K, V]) Iter() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		sm.mut.RLock()
		defer sm.mut.RUnlock()
		for k, v := range sm.values {
			if !yield(k, v) {
				return
			}
		}
	}
}

func (sm *Sync[K, V]) Len() (i int) {
	sm.mut.RLock()
	i = len(sm.values)
	sm.mut.RUnlock()
	return
}

// Sync is a concurrency-safe hash map wrapper.
type Sync[K comparable, V any] struct {
	values map[K]V
	mut    sync.RWMutex
}

// NewSync initializes and returns a new Sync.
func NewSync[K comparable, V any]() *Sync[K, V] { return &Sync[K, V]{values: make(map[K]V)} }

// Get retrieves a value by key. It returns the value and a boolean indicating if the key exists.
func (sm *Sync[K, V]) Get(k K) (res V, ok bool) {
	sm.mut.RLock()
	res, ok = sm.values[k]
	sm.mut.RUnlock()
	return
}

// Set adds or updates a value in the map for a given key.
func (sm *Sync[K, V]) Set(k K, v V) {
	sm.mut.Lock()
	sm.values[k] = v
	sm.mut.Unlock()
}

// Del removes a value by key. It returns a boolean indicating if the key was found and removed.
func (sm *Sync[K, V]) Del(k K) (ok bool) {
	sm.mut.Lock()
	if _, ok = sm.values[k]; ok {
		delete(sm.values, k)
	}
	sm.mut.Unlock()
	return
}

func (sm *Sync[K, V]) lockless_get(k K) (res V, ok bool) {
	res, ok = sm.values[k]
	return
}
func (sm *Sync[K, V]) lockless_set(k K, v V) V { sm.values[k] = v; return v }
func (sm *Sync[K, V]) lockless_del(k K) (ok bool) {
	if _, ok = sm.values[k]; ok {
		delete(sm.values, k)
	}
	return
}
