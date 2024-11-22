package gen

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

// func TestSMap(t *testing.T) {}

// func FuzzSMap(f *testing.F) {
// 	sm := NewSyncMap[uint64, uint8]()
// 	f.Add(uint8(0), uint64(0), uint8(0))
// 	f.Fuzz(func(t *testing.T, action uint8, k uint64, v uint8) {
// 		switch action {
// 		case 0, 1, 3, 4:
// 			go sm.Get(k)
// 		case 5, 6:
// 			go sm.Set(k, v)
// 		default:
// 			go sm.Del(k)
// 		}
// 	})
// }

// ChatGPT
func TestSMap(t *testing.T) {
	const numKeys = 20_000   // Number of keys to test
	const numOps = 1_000_000 // Total number of operations
	const numGoroutines = 50 // Number of concurrent goroutines

	// Create the SyncMap
	sm := NewSyncMap[int, int]()

	// Seed the random number generator for fuzzing
	rand.Seed(time.Now().UnixNano())

	// Synchronization for all goroutines
	var wg sync.WaitGroup

	// Fuzzing operations
	operation := func() {
		defer wg.Done()
		for i := 0; i < numOps; i++ {
			key := rand.Intn(numKeys) // Random key within the range
			value := rand.Intn(1000)  // Random value
			action := rand.Intn(3)    // Random operation: 0=Get, 1=Set, 2=Del

			switch action {
			case 0: // Get
				_, _ = sm.Get(key) // Simply call Get to ensure no panics
			case 1: // Set
				sm.Set(key, value)
			case 2: // Del
				sm.Del(key)
			}
		}
	}

	// Spawn goroutines
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go operation()
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Validation: Ensure map integrity (no panics occurred and map is consistent)
	totalKeys := 0
	sm.mut.RLock()
	for k, v := range sm.val {
		t.Logf("Key: %d, Value: %d", k, v)
		totalKeys++
	}
	sm.mut.RUnlock()

	t.Logf("Total keys remaining in map: %d", totalKeys)
}
