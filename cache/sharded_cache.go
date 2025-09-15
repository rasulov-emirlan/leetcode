package cache

import (
	"time"
)

type Hasher[K comparable] func(K) uint64

type ShardedCache[K comparable, T any] struct {
	shards []*Cache[K, T]
	mask   uint64 // len(shards)-1, power of two
	hash   Hasher[K]
}

// round up to the next power of two
func nextPow2(n int) int {
	if n < 1 {
		return 1
	}
	p := 1
	for p < n {
		p <<= 1
	}
	return p
}

func NewShardedCache[K comparable, T any](numShards int, sweepInterval time.Duration, hasher Hasher[K]) *ShardedCache[K, T] {
	n := nextPow2(numShards)
	shards := make([]*Cache[K, T], n)
	for i := 0; i < n; i++ {
		shards[i] = NewCache[K, T](sweepInterval) // your per-shard cache constructor with janitor
	}
	return &ShardedCache[K, T]{
		shards: shards,
		mask:   uint64(n - 1),
		hash:   hasher,
	}
}

func (sc *ShardedCache[K, T]) shardFor(key K) *Cache[K, T] {
	idx := sc.hash(key) & sc.mask
	return sc.shards[idx]
}

func (sc *ShardedCache[K, T]) Set(key K, value T, ttl time.Duration) {
	sc.shardFor(key).Set(key, value, ttl)
}

func (sc *ShardedCache[K, T]) Get(key K) (T, bool) {
	return sc.shardFor(key).Get(key)
}

func (sc *ShardedCache[K, T]) Delete(key K) {
	sc.shardFor(key).Delete(key)
}

func (sc *ShardedCache[K, T]) Len() int {
	total := 0
	for _, s := range sc.shards {
		total += s.Len()
	}
	return total
}

func (sc *ShardedCache[K, T]) Close() {
	for _, s := range sc.shards {
		s.Close()
	}
}

// ---------- Example hashers (fast, no allocations) ----------

func StringHasher(s string) uint64 {
	// FNV-1a 64
	var h uint64 = 1469598103934665603
	const prime = 1099511628211
	for i := 0; i < len(s); i++ {
		h ^= uint64(s[i])
		h *= prime
	}
	return h
}

func BytesHasher(b []byte) uint64 {
	var h uint64 = 1469598103934665603
	const prime = 1099511628211
	for i := 0; i < len(b); i++ {
		h ^= uint64(b[i])
		h *= prime
	}
	return h
}

func Uint64Hasher(u uint64) uint64 { return u }
