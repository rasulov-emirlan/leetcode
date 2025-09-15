package cache

import (
	"testing"
	"time"
)

func TestShardedCache(t *testing.T) {
	sc := NewShardedCache[string, string](4, 50*time.Millisecond, StringHasher)
	sc.Set("key1", "value1", 200*time.Millisecond)
	val, ok := sc.Get("key1")
	if !ok || val != "value1" {
		t.Error("Expected to retrieve value1")
	}

	time.Sleep(250 * time.Millisecond)
	_, ok = sc.Get("key1")
	if ok {
		t.Error("Expected cache to be expired")
	}

	sc.Set("key2", "value2", 500*time.Millisecond)
	val, ok = sc.Get("key2")
	if !ok || val != "value2" {
		t.Error("Expected to retrieve value2")
	}

	sc.Delete("key2")
	_, ok = sc.Get("key2")
	if ok {
		t.Error("Expected key2 to be deleted")
	}
}

func TestShardedCacheConcurrently(t *testing.T) {
	sc := NewShardedCache[string, string](4, 50*time.Millisecond, StringHasher)
	numGoroutines := 10
	numOperations := 100

	// Writer goroutines
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			for j := 0; j < numOperations; j++ {
				key := "key" + string(rune(id)) + "_" + string(rune(j))
				sc.Set(key, "value"+string(rune(j)), 200*time.Millisecond)
				time.Sleep(5 * time.Millisecond)
			}
		}(i)
	}

	// Reader goroutines
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			for j := 0; j < numOperations; j++ {
				key := "key" + string(rune(id)) + "_" + string(rune(j))
				sc.Get(key)
				time.Sleep(5 * time.Millisecond)
			}
		}(i)
	}

	// Wait for a while to let goroutines finish
	time.Sleep(3 * time.Second)
}
