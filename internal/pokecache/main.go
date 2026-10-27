// Package pokecache
package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entry map[string]cacheEntry
	mu    sync.Mutex
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{entry: make(map[string]cacheEntry)}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	var entry cacheEntry
	entry.createdAt = time.Now()
	entry.val = val
	c.entry[key] = entry
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	entry, ok := c.entry[key]
	c.mu.Unlock()
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for key, value := range c.entry {
			t := time.Now()
			elapsed := t.Sub(value.createdAt)
			if elapsed > interval {
				delete(c.entry, key)
			}
		}
		c.mu.Unlock()
	}
}
