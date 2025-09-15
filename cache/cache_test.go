package cache

import (
	"sync"
	"testing"
	"time"
)

func TestCacheConcurrently(t *testing.T) {
	c := NewCache[string, string](50 * time.Millisecond)
	var wg sync.WaitGroup
	numGoroutines := 10
	numOperations := 100

	// Writer goroutines
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				key := "key" + string(rune(id)) + "_" + string(rune(j))
				c.Set(key, "value"+string(rune(j)), 200*time.Millisecond)
				time.Sleep(5 * time.Millisecond)
			}
		}(i)
	}

	// Reader goroutines
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numOperations; j++ {
				key := "key" + string(rune(id)) + "_" + string(rune(j))
				c.Get(key)
				time.Sleep(5 * time.Millisecond)
			}
		}(i)
	}

	wg.Wait()
}

func TestCacheExpiration(t *testing.T) {
	c := NewCache[string, string](50 * time.Millisecond)
	c.Set("temp", "value", 100*time.Millisecond)

	time.Sleep(150 * time.Millisecond)

	_, ok := c.Get("temp")
	if ok {
		t.Error("Expected cache to be expired")
	}
}

func BenchmarkCacheSet(b *testing.B) {
	c := NewCache[int, int](100 * time.Millisecond)
	for i := 0; i < b.N; i++ {
		c.Set(i, i*2, 1*time.Second)
	}
}

func BenchmarkCacheGet(b *testing.B) {
	c := NewCache[int, int](100 * time.Millisecond)
	for i := 0; i < 1000; i++ {
		c.Set(i, i*2, 1*time.Second)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Get(i % 1000)
	}
}
