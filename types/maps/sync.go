package maps

import (
	"iter"
	"sync"
)

type Map[K comparable, V any] interface {
	Get(key K) (val V, ok bool)
	Set(key K, val V) (ok bool)
	Del(keys ...K)
	Iter() iter.Seq2[K, V]
	GetFull(key K) (val V, ok bool, valid bool)
}

func New[K comparable, V any](isValid func(K, V) bool) Map[K, V] {
	base := &Sync[K, V]{
		values: make(map[K]V),
	}
	switch isValid {
	case nil:
		return base
	default:
		return &Validated[K, V]{base, isValid}
	}
}

// Sync is a thread safe map wrapper.
type Sync[K comparable, V any] struct {
	values map[K]V
	mut    sync.RWMutex
}

func (sm *Sync[K, V]) Get(key K) (val V, ok bool) {
	sm.mut.RLock()
	val, ok = sm.values[key]
	sm.mut.RUnlock()
	return
}

func (sm *Sync[K, V]) GetFull(key K) (val V, ok bool, valid bool) {
	valid = true
	val, ok = sm.Get(key)
	return
}

func (sm *Sync[K, V]) Set(key K, val V) (ok bool) {
	sm.mut.Lock()
	sm.values[key] = val
	sm.mut.Unlock()
	return true
}

func (sm *Sync[K, V]) Del(keys ...K) {
	sm.mut.Lock()
	for _, key := range keys {
		delete(sm.values, key)
	}
	sm.mut.Unlock()
}

func (sm *Sync[K, V]) Len() (i int) {
	sm.mut.RLock()
	i = len(sm.values)
	sm.mut.RUnlock()
	return
}

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

type Validated[K comparable, V any] struct {
	*Sync[K, V]
	isValid func(K, V) bool
}

func (em *Validated[K, V]) Get(key K) (res V, ok bool) {
	res, ok = em.Sync.Get(key)
	if !ok {
		return
	}
	if !em.isValid(key, res) {
		em.Del(key)
	}
	return
}

func (em *Validated[K, V]) GetFull(key K) (res V, ok bool, valid bool) {
	valid = true
	res, ok = em.Sync.Get(key)
	if !ok {
		return
	}
	if !em.isValid(key, res) {
		em.Del(key)
		valid = false
	}
	return
}

func (em *Validated[K, V]) Set(key K, val V) (ok bool) {
	if !em.isValid(key, val) {
		return
	}
	em.Sync.Set(key, val)
	return true
}
func (em *Validated[K, V]) Iter() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		invalidKeys := []K{}
		em.mut.RLock()
		for k, val := range em.values {
			if !em.isValid(k, val) {
				invalidKeys = append(invalidKeys, k)
				continue
			}
			if !yield(k, val) {
				break
			}
		}
		em.mut.RUnlock()
		if len(invalidKeys) == 0 {
			return
		}
		em.mut.Lock()
		for _, val := range invalidKeys {
			delete(em.values, val)
		}
		em.mut.Unlock()
	}
}
