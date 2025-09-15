package cache

import (
	"sync"
	"time"
)

type iterm[T any] struct {
	value      T
	expiration time.Time
}

type Cache[K comparable, T any] struct {
	m           map[K]iterm[T]
	lock        sync.RWMutex
	closeCh     chan struct{}
	isDone      bool
	whenToSweep time.Duration
}

func NewCache[K comparable, T any](sweepInterval time.Duration) *Cache[K, T] {
	c := &Cache[K, T]{
		m:           make(map[K]iterm[T]),
		lock:        sync.RWMutex{},
		closeCh:     make(chan struct{}),
		isDone:      false,
		whenToSweep: sweepInterval,
	}
	go c.janitor()
	return c
}

func (c *Cache[K, T]) Set(key K, value T, ttl time.Duration) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.m[key] = iterm[T]{
		value:      value,
		expiration: time.Now().Add(ttl),
	}
}

func (c *Cache[K, T]) Get(key K) (T, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()
	item, exists := c.m[key]
	if !exists || time.Now().After(item.expiration) {
		var zero T
		return zero, false
	}
	return item.value, true
}

func (c *Cache[K, T]) Delete(key K) {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.m, key)
}

func (c *Cache[K, T]) janitor() {
	ticker := time.NewTicker(c.whenToSweep)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.sweep()
		case <-c.closeCh:
			return
		}
	}
}

func (c *Cache[K, T]) sweep() {
	c.lock.Lock()
	defer c.lock.Unlock()

	now := time.Now()
	for key, item := range c.m {
		if now.After(item.expiration) {
			delete(c.m, key)
		}
	}
}

func (c *Cache[K, T]) Close() {
	c.lock.Lock()
	defer c.lock.Unlock()
	if !c.isDone {
		close(c.closeCh)
		c.isDone = true
	}
}

func (c *Cache[K, T]) Len() int {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return len(c.m)
}
