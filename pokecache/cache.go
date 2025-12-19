package pokecache

import (
	"sync"
	"time"
)

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	cache := cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
	c.cache[key] = cache
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	if v, ok := c.cache[key]; ok {
		return v.val, true
	}

	return []byte{}, false
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		currentTime := time.Now().UTC()
		c.mu.Lock()
		for k, e := range c.cache {
			if e.createdAt.Before(currentTime.Add(-interval)) {
				delete(c.cache, k)
			}
		}
		c.mu.Unlock()
	}
}

func NewCache(t time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mu:    &sync.Mutex{},
	}
	go c.reapLoop(t + 15*time.Second)

	return c
}
