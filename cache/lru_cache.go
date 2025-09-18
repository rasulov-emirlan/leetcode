package cache

import "sync"

type lruItem[K comparable, T any] struct {
	Key   K
	Value T
	Next  *lruItem[K, T]
	Prev  *lruItem[K, T]
}

type LRUCache[K comparable, T any] struct {
	head *lruItem[K, T]
	tail *lruItem[K, T]
	m    map[K]*lruItem[K, T]
	lock sync.RWMutex
	size uint32
}

func NewLRUCache[K comparable, T any](size uint32) *LRUCache[K, T] {
	if size == 0 {
		size = 1
	}
	return &LRUCache[K, T]{
		m:    make(map[K]*lruItem[K, T], size),
		size: size,
	}
}

func (l *LRUCache[K, T]) Get(key K) (T, bool) {
	l.lock.RLock()
	defer l.lock.RUnlock()

	if item, ok := l.m[key]; ok {
		l.moveToFront(item)
		return item.Value, true
	}
	var zero T
	return zero, false
}

func (l *LRUCache[K, T]) Set(key K, value T) {
	l.lock.Lock()
	defer l.lock.Unlock()

	if item, ok := l.m[key]; ok {
		item.Value = value
		l.moveToFront(item)
		return
	}

	newItem := &lruItem[K, T]{Key: key, Value: value}
	l.m[key] = newItem
	l.addToFront(newItem)

	if uint32(len(l.m)) > l.size {
		l.removeOldest()
	}
}

func (l *LRUCache[K, T]) Delete(key K) {
	l.lock.Lock()
	defer l.lock.Unlock()

	if item, ok := l.m[key]; ok {
		l.removeItem(item)
		delete(l.m, key)
	}
}

func (l *LRUCache[K, T]) Len() int {
	l.lock.RLock()
	defer l.lock.RUnlock()
	return len(l.m)
}

func (l *LRUCache[K, T]) moveToFront(item *lruItem[K, T]) {
	if item == l.head {
		return
	}
	l.removeItem(item)
	l.addToFront(item)
}

func (l *LRUCache[K, T]) addToFront(item *lruItem[K, T]) {
	item.Next = l.head
	item.Prev = nil
	if l.head != nil {
		l.head.Prev = item
	}
	l.head = item
	if l.tail == nil {
		l.tail = item
	}
}

func (l *LRUCache[K, T]) removeItem(item *lruItem[K, T]) {
	if item.Prev != nil {
		item.Prev.Next = item.Next
	} else {
		l.head = item.Next
	}
	if item.Next != nil {
		item.Next.Prev = item.Prev
	} else {
		l.tail = item.Prev
	}
}

func (l *LRUCache[K, T]) removeOldest() {
	if l.tail != nil {
		delete(l.m, l.tail.Key)
		l.removeItem(l.tail)
	}
}
