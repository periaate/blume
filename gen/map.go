package gen

import "sync"

// SyncMap is a concurrency-safe hash map wrapper.
type SyncMap[K comparable, V any] struct {
	val map[K]V
	mut sync.RWMutex
}

// NewSyncMap initializes and returns a new SyncMap.
func NewSyncMap[K comparable, V any]() *SyncMap[K, V] {
	return &SyncMap[K, V]{val: make(map[K]V)}
}

// Get retrieves a value by key. It returns the value and a boolean indicating if the key exists.
func (sm *SyncMap[K, V]) Get(k K) (V, bool) {
	sm.mut.RLock()
	defer sm.mut.RUnlock()
	v, ok := sm.val[k]
	return v, ok
}

// Set adds or updates a value in the map for a given key.
func (sm *SyncMap[K, V]) Set(k K, v V) {
	sm.mut.Lock()
	defer sm.mut.Unlock()
	sm.val[k] = v
}

// Del removes a value by key. It returns a boolean indicating if the key was found and removed.
func (sm *SyncMap[K, V]) Del(k K) bool {
	sm.mut.Lock()
	defer sm.mut.Unlock()
	if _, ok := sm.val[k]; ok {
		delete(sm.val, k)
		return true
	}
	return false
}
