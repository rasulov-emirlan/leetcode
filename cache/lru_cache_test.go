package cache

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {
	lru := NewLRUCache[string, string](2)

	lru.Set("key1", "value1")
	lru.Set("key2", "value2")
	if val, ok := lru.Get("key1"); !ok || val != "value1" {
		t.Error("Expected to retrieve value1")
	}

	lru.Set("key3", "value3") // This should evict key2
	if _, ok := lru.Get("key2"); ok {
		t.Error("Expected key2 to be evicted")
	}

	lru.Delete("key1")
	if _, ok := lru.Get("key1"); ok {
		t.Error("Expected key1 to be deleted")
	}

	if length := lru.Len(); length != 1 {
		t.Errorf("Expected cache length to be 1, got %d", length)
	}
}

func BenchmarkLRUCache(b *testing.B) {
	lru := NewLRUCache[string, string](1000)
	for i := 0; i < b.N; i++ {
		lru.Set(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
}
