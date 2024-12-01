package maps

import (
	"fmt"
	"iter"
	"sync"
)

func (sm *Sync[K, V]) Keys() iter.Seq[K] {
	return func(yield func(K) bool) {
		sm.mut.RLock()
		defer sm.mut.RUnlock()
		for k := range sm.Values {
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
		for _, v := range sm.Values {
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
		for k, v := range sm.Values {
			if !yield(k, v) {
				return
			}
		}
	}
}

func (sm *Sync[K, V]) Len() int {
	sm.mut.RLock()
	defer sm.mut.RUnlock()
	return len(sm.Values)
}

// Sync is a concurrency-safe hash map wrapper.
type Sync[K comparable, V any] struct {
	Values map[K]V
	mut    sync.RWMutex

	Hooks Hooks[K, V]
}

// NewSync initializes and returns a new Sync.
func NewSync[K comparable, V any](hooks ...Hooks[K, V]) *Sync[K, V] {
	var hook Hooks[K, V]
	if len(hooks) == 0 {
		hook = Hooks[K, V]{}
	} else {
		hook = hooks[0]
	}

	return &Sync[K, V]{
		Values: make(map[K]V),
		Hooks:  hook,
	}
}

// Get retrieves a value by key. It returns the value and a boolean indicating if the key exists.
func (sm *Sync[K, V]) Get(k K) (res V, ok bool) {
	sm.mut.RLock()
	res, ok = sm.Values[k]
	sm.mut.RUnlock()
	if sm.Hooks.Get == nil {
		return
	}

	ak, av, op := sm.Hooks.Get(k, res)
	switch op {
	case OP_DEL:
		ok = sm.Del(ak)
	case OP_SET:
		ok = sm.Set(ak, av)
		res = av
	}
	return
}

// Set adds or updates a value in the map for a given key.
func (sm *Sync[K, V]) Set(k K, v V) (ok bool) {
	fmt.Println("Set A")
	if sm.Hooks.Set != nil {
		fmt.Println("Set B")
		ak, av, op := sm.Hooks.Set(k, v)
		fmt.Println("Set C")
		switch op {
		case OP_SET:
			fmt.Println("Set D")
			k, v = ak, av
		case OP_DEL:
			fmt.Println("Set E")
			ok = sm.Del(ak)
			return
		case OP_RET:
			return
		}
	}

	fmt.Println("Set F")
	sm.mut.Lock()
	fmt.Println("Set G")
	defer sm.mut.Unlock()
	fmt.Println("Set H")
	sm.Values[k] = v
	return true
}

// Del removes a value by key. It returns a boolean indicating if the key was found and removed.
func (sm *Sync[K, V]) Del(k K) (ok bool) {
	if sm.Hooks.Del != nil {
		sm.Hooks.Del(k)
	}

	sm.mut.Lock()
	defer sm.mut.Unlock()
	if _, ok := sm.Values[k]; ok {
		delete(sm.Values, k)
		return true
	}
	return false
}
