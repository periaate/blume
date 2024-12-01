package maps

import (
	"testing"
	"time"
)

// TestGetExpired ensures that expired items are removed when accessed.
func TestGetExpired(t *testing.T) {
	m := NewExpiring[string, string]()

	key := "key"
	value := "value"

	// Set an item with an expiration in the past.
	m.Set(key, value, time.Now().Add(-time.Second))

	// Try to get the expired item.
	_, ok := m.Get(key)
	if ok {
		t.Fatalf("expected no value for key %q as it should be expired", key)
	}

	// Ensure the item was removed.
	_, ok = m.Sync.Get(key)
	if ok {
		t.Fatalf("expected expired item to be deleted from the map")
	}
}

// TestGetValid ensures that valid (not expired) items are retrievable.
func TestGetValid(t *testing.T) {
	m := NewExpiring[string, string]()

	key := "key"
	value := "value"

	// Set an item with a future expiration time.
	m.Set(key, value, time.Now().Add(time.Second*5))

	// Try to get the valid item.
	res, ok := m.Get(key)
	if !ok {
		t.Fatalf("expected value for key %q", key)
	}
	if res != value {
		t.Fatalf("expected value %q, got %q", value, res)
	}
}

// TestSetExpired ensures that setting an already expired item does not store it.
func TestSetExpired(t *testing.T) {
	m := NewExpiring[string, string]()

	key := "key"
	value := "value"

	// Try to set an item with an expiration in the past.
	ok := m.Set(key, value, time.Now().Add(-time.Second))
	if ok {
		t.Fatalf("expected Set to return false for an expired item")
	}

	// Ensure the item was not added.
	_, exists := m.Sync.Get(key)
	if exists {
		t.Fatalf("expected no value for key %q", key)
	}
}

// TestDel ensures that Del correctly removes items.
func TestDel(t *testing.T) {
	m := NewExpiring[string, string]()

	key := "key"
	value := "value"

	// Set an item with a future expiration time.
	m.Set(key, value, time.Now().Add(time.Second*5))

	// Delete the item.
	ok := m.Del(key)
	if !ok {
		t.Fatalf("expected Del to return true for existing key %q", key)
	}

	// Ensure the item was removed.
	_, exists := m.Sync.Get(key)
	if exists {
		t.Fatalf("expected no value for key %q after deletion", key)
	}
}

// TestConcurrentAccess ensures thread-safety for concurrent Get and Set operations.
func TestConcurrentAccess(t *testing.T) {
	m := NewExpiring[string, string]()

	key := "key"
	value := "value"

	// Use a channel to synchronize goroutines.
	done := make(chan struct{})

	// Concurrent setter.
	go func() {
		exp := time.Now().Add(time.Millisecond * 10)
		for i := 0; i < 10_000; i++ {
			if !m.Set(key, value, exp) {
				t.Fatalf("set returned non true")
			}
		}
		done <- struct{}{}
	}()

	// Concurrent deleter.
	go func() {
		for i := 0; i < 10_000; i++ {
			m.Del(key)
		}
		done <- struct{}{}
	}()

	// Concurrent getter.
	go func() {
		for i := 0; i < 10_000; i++ {
			m.Get(key)
		}
		done <- struct{}{}
	}()

	// Wait for both goroutines to complete.
	<-done
	<-done
	<-done
}
