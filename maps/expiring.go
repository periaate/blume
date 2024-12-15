package maps

import (
	"iter"
	"sync"
	"time"

	. "github.com/periaate/blume"
)

type ExpItem[V any] struct {
	Value   V
	Expires time.Time
}

// Expiring is a thread safe map where values have expiration dates.
// Expiring does not automatically clear expired items, rather, they are deleted on Get.
type Expiring[K comparable, V any] struct {
	*Sync[K, ExpItem[V]]
	del_ch chan(K)
	cls_ch chan(chan(any))
}

// honestly, no clue how this works, but it seems to? prod ready if you ask me :^)
func expiration_worker[K comparable, V any](exp *Expiring[K, V]) (chan(K), chan(chan(any))) {
	del_ch := make(chan(K))
	cls_ch := make(chan(chan(any)))
	clearing := false // avoids deadlock with queueing deletions
	q := head[K]{}

	mut := sync.Mutex{}
	
	go func() {
		for {
			select {
			case k := <- del_ch:
				mut.Lock()
				q.Push(k, Top)
				mut.Unlock()
				if !clearing && q.Len >= 1000 { clearing = true; cls_ch <- nil }
			}
		}
	}()

	go func() {
		for {
			fin_ch := <- cls_ch
			mut.Lock()
			exp.mut.Lock()
			clearing = true
			for {
				opt := q.Pop(Bot)
				if !opt.Ok {
					if q.Len == 0 { break }
					continue
				}
				exp.lockless_del(opt.Value)
			}
			clearing = false
			exp.mut.Unlock()
			mut.Unlock()
			if fin_ch != nil { fin_ch <- nil }
		}
	}()

	return del_ch, cls_ch
}

// Iter returns a sequence of key-value pairs in the map.
func (em *Expiring[K, V]) Iter() iter.Seq2[K, V] {
	return func(yield func(K, V) bool) {
		em.mut.RLock()
		defer em.mut.RUnlock()
		for k, v := range em.values {
			if isExpired(v.Expires) { em.del_ch <- k; continue }
			if !yield(k, v.Value) { return }
		}
	}
}

// NewSync initializes and returns a new Sync.
func NewExpiring[K comparable, V any]() *Expiring[K, V] {
	exp := &Expiring[K, V]{Sync: NewSync[K, ExpItem[V]]()}
	exp.del_ch, exp.cls_ch = expiration_worker[K, V](exp)
	return exp
}

func (em *Expiring[K, V]) Flush() {
	fin_ch := make(chan(any))
	em.cls_ch <- fin_ch
	<- fin_ch
}

func isExpired(t time.Time) bool { return t.Before(time.Now()) }

// Get retrieves a value by key. It returns the value and a boolean indicating if the key exists.
func (em *Expiring[K, V]) Get(k K) Option[V] {
	it := em.Sync.Get(k)
	if !it.Ok { return None[V]() }
	if isExpired(it.Value.Expires) {
		em.del_ch <- k
		return None[V]()
	}
	return Some(it.Value.Value)
}

// Set adds or updates a value in the map for a given key.
func (em *Expiring[K, V]) Set(k K, v V, dur time.Duration) Option[V] {
	expires := time.Now().Add(dur)
	if isExpired(expires) { return None[V]() }
	return Some(em.Sync.Set(k, ExpItem[V]{v, expires}).Value)
}

func (em *Expiring[K, V]) Del(k K) { em.del_ch <- k }
