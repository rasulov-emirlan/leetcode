package lrucache

type LRUCache struct {
	m    map[int]*item
	head *item
	tail *item
	cap  int
}

type item struct {
	key  int
	val  int
	prev *item
	next *item
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		m:   make(map[int]*item),
		cap: capacity,
	}
}

func (lru *LRUCache) Get(key int) int {
	v, ok := lru.m[key]
	if !ok {
		return -1
	}

	lru.moveToHead(v)
	return v.val
}

func (lru *LRUCache) Put(key, value int) {
	// If key exists, update value and move to head
	if existing, ok := lru.m[key]; ok {
		existing.val = value
		lru.moveToHead(existing)
		return
	}

	i := &item{val: value, key: key}
	lru.attachToHead(i)
	lru.m[key] = i

	if len(lru.m) > lru.cap {
		lru.deleteBottom()
	}
}

func (lru *LRUCache) moveToHead(i *item) {
	if i == lru.head {
		return
	}
	// If node is already in the list, detach it first
	if i == lru.tail || i.prev != nil || i.next != nil {
		lru.detach(i)
	}
	lru.attachToHead(i)
}

func (lru *LRUCache) attachToHead(i *item) {
	if lru.head == nil { // empty list
		lru.head = i
		lru.tail = i
		return
	}
	i.prev = lru.head
	i.next = nil
	lru.head.next = i
	lru.head = i
}

func (lru *LRUCache) detach(i *item) {
	prev := i.prev
	next := i.next
	if prev != nil {
		prev.next = next
	} else {
		// i was tail
		lru.tail = next
	}
	if next != nil {
		next.prev = prev
	} else {
		// i was head
		lru.head = prev
	}
	i.prev = nil
	i.next = nil
}

func (lru *LRUCache) deleteBottom() {
	if lru.tail == nil {
		return
	}
	toRemove := lru.tail
	lru.detach(toRemove)
	delete(lru.m, toRemove.key)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
