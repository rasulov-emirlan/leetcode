package lrucache

import "testing"

func TestLRUCache(t *testing.T) {
	// ["LRUCache","put","put","get","put","get","put","get","get","get"]
	// [[2],[1,1],[2,2],[1],[3,3],[2],[4,4],[1],[3],[4]]

	// expect: [null,null,null,1,null,-1,null,-1,3,4]

	lru := Constructor(2)
	lru.Put(1, 1)
	lru.Put(2, 2)

	if got := lru.Get(1); got != 1 {
		t.Errorf("Get(1) = %v, want 1", got)
	}

	lru.Put(3, 3)

	if got := lru.Get(2); got != -1 {
		t.Errorf("Get(2) = %v, want -1", got)
	}

	lru.Put(4, 4)

	if got := lru.Get(1); got != -1 {
		t.Errorf("Get(1) = %v, want -1", got)
	}

	if got := lru.Get(3); got != 3 {
		t.Errorf("Get(3) = %v, want 3", got)
	}

	if got := lru.Get(4); got != 4 {
		t.Errorf("Get(4) = %v, want 4", got)
	}
}

func BenchmarkLRUCache(b *testing.B) {
	lru := Constructor(1000)
	for i := 0; i < b.N; i++ {
		lru.Put(i, i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		lru.Get(i)
	}
}
