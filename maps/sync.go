package maps

import (
	"iter"
	"sync"
)

// Sync is a concurrency-safe hash map wrapper.
type Sync[K comparable, V any] struct {
	val map[K]V
	mut sync.RWMutex
}

func (sm *Sync[K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		sm.mut.RLock()
		defer sm.mut.RUnlock()
		for k := range sm.val {
			if !yield(k) {
				return
			}
		}
	}
}

func (sm *Sync[K, V]) Vals() iter.Seq[V] {
	return func(yield func(V) bool) {
		sm.mut.RLock()
		defer sm.mut.RUnlock()
		for _, v := range sm.val {
			if !yield(v) {
				return
			}
		}
	}
}

func (sm *Sync[K, V]) Iter() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		sm.mut.RLock()
		defer sm.mut.RUnlock()
		for k, v := range sm.val {
			if !yield(k, v) {
				return
			}
		}
	}
}

func (sm *Sync[K, V]) Len() int {
	sm.mut.RLock()
	defer sm.mut.RUnlock()
	return len(sm.val)
}

// NewSync initializes and returns a new Sync.
func NewSync[K comparable, V any]() *Sync[K, V] {
	return &Sync[K, V]{val: make(map[K]V)}
}

// Get retrieves a value by key. It returns the value and a boolean indicating if the key exists.
func (sm *Sync[K, V]) Get(k K) (res V, ok bool) {
	sm.mut.RLock()
	defer sm.mut.RUnlock()
	res, ok = sm.val[k]
	return
}

// Set adds or updates a value in the map for a given key. This always returns true.
func (sm *Sync[K, V]) Set(k K, v V) (ok bool) {
	sm.mut.Lock()
	defer sm.mut.Unlock()
	sm.val[k] = v
	return true
}

// Del removes a value by key. It returns a boolean indicating if the key was found and removed.
func (sm *Sync[K, V]) Del(k K) (ok bool) {
	sm.mut.Lock()
	defer sm.mut.Unlock()
	if _, ok := sm.val[k]; ok {
		delete(sm.val, k)
		return true
	}
	return false
}
