package maps

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func newKVP(sec int) (string, string, time.Duration) {
	k := fmt.Sprintf("key[%d]", sec)
	v := fmt.Sprintf("value[%d]", sec)
	return k, v, time.Second * time.Duration(sec)
}

// TestGetExpired ensures that expired items are removed when accessed.
func TestGetExpired(t *testing.T) {
	m := NewExpiring[string, string]()
	k, v, e := newKVP(-1)
	assert.False(t, m.Set(k, v, e).Ok)
	assert.False(t, m.Get(k).Ok)
}

// TestGetValid ensures that valid (not expired) items are retrievable.
func TestGetValid(t *testing.T) {
	m := NewExpiring[string, string]()
	k, v, e := newKVP(1)
	assert.True(t, m.Set(k, v, e).Ok)
	assert.True(t, m.Get(k).Ok)
	assert.Equal(t, v, m.Get(k).Value)
}

// TestDel ensures that Del correctly removes items.
func testDel(m *Expiring[string, string], Range [2]int, t *testing.T) {
	for i := range Range[1] - Range[0] {
		k, v, e := newKVP(i+1 + Range[0])
		assert.Truef(t, m.Set(k, v, e).Ok, "tried to set, but couldn't %s", k)
	}

	for i := range Range[1] - Range[0] {
		k, _, _ := newKVP(i)
		m.Del(k)
	}
}

func TestDel(t *testing.T) {
	length := 1000
	m := NewExpiring[string, string]()
	testDel(m, [2]int{0, length}, t)

	for i := range length+1 {
		k, _, _ := newKVP(i)
		m.Del(k)
	}

	for i := range length+1 {
		k, _, _ := newKVP(i)
		if !assert.False(t, m.Get(k).Ok) { fmt.Println("Key:", k) }
	}
}

func TestConcurrentAccess(t *testing.T) {
	count := 10
	length := 1000
	m := NewExpiring[string, string]()
	wg := sync.WaitGroup{}
	for i := range count {
		wg.Add(1)
		from := i * length
		go func() {
			testDel(m, [2]int{from, from + length}, t)
			wg.Done()
		}()
	}
	wg.Wait()

	for i := range count*length+1 {
		k, _, _ := newKVP(i)
		m.Del(k)
	}

	for i := range count*length {
		k, _, _ := newKVP(i)
		if !assert.False(t, m.Get(k).Ok) { fmt.Println("Key:", k) }
	}
}
