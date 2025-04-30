package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt	time.Time
	val			[]byte
}

type Cache struct {
	cache	map[string]cacheEntry
	mu		*sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		cache:	make(map[string]cacheEntry),
		mu:		&sync.Mutex{},
	}
	go cache.reapLoop(interval)

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = cacheEntry{
		createdAt: 	time.Now().UTC(),
		val: 		val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	
	entry, ok := c.cache[key]
	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, interval time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-interval)) {
			delete(c.cache, k)
		}
	}
}