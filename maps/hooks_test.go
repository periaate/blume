package maps

import (
	"sync/atomic"
	"testing"
)

func TestSyncHooks(t *testing.T) {
	type testKey string
	type testValue string

	var getCount, setCount, delCount int32

	hooks := Hooks[testKey, testValue]{
		Get: func(key testKey, value testValue) (testKey, testValue, Operation) {
			atomic.AddInt32(&getCount, 1)
			if key == "hooked" {
				return key, "hooked_value", OP_SET
			}
			return key, value, OP_RET
		},
		Set: func(key testKey, value testValue) (testKey, testValue, Operation) {
			atomic.AddInt32(&setCount, 1)
			if key == "delete_me" {
				return key, value, OP_DEL
			}
			return key, value, OP_SET
		},
		Del: func(key testKey) (testKey, Operation) {
			atomic.AddInt32(&delCount, 1)
			return key, OP_NIL
		},
	}

	m := NewSync(hooks)

	// Test Set Hook
	m.Set("normal_key", "normal_value")
	m.Set("delete_me", "value_to_delete")

	if count := atomic.LoadInt32(&setCount); count != 2 {
		t.Errorf("Expected 2 calls to Set hook, got %d", count)
	}

	if _, ok := m.Get("delete_me"); ok {
		t.Error("Key 'delete_me' should have been deleted by the Set hook")
	}

	// Test Get Hook
	m.Set("hooked", "original_value")
	value, _ := m.Get("hooked")

	if value != "hooked_value" {
		t.Errorf("Expected Get hook to modify value to 'hooked_value', got %v", value)
	}

	if count := atomic.LoadInt32(&getCount); count != 2 {
		t.Errorf("Expected 1 call to Get hook, got %d", count)
	}

	// Test Del Hook
	m.Set("delete_me", "value_to_delete")
	m.Del("delete_me")

	if count := atomic.LoadInt32(&delCount); count != 3 {
		t.Errorf("Expected 1 call to Del hook, got %d", count)
	}

	if _, ok := m.Get("delete_me"); ok {
		t.Error("Key 'delete_me' should have been deleted")
	}
}

func TestSyncWithoutHooks(t *testing.T) {
	m := NewSync[string, string]()

	// Basic Set and Get without hooks
	m.Set("key", "value")
	val, ok := m.Get("key")

	if !ok || val != "value" {
		t.Errorf("Expected 'value' for 'key', got %v (exists: %v)", val, ok)
	}

	// Delete without hooks
	m.Del("key")
	_, ok = m.Get("key")
	if ok {
		t.Error("Expected 'key' to be deleted")
	}
}
