package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{
		entries: make(map[string]cacheEntry),
		// missed mu
		mu:       &sync.Mutex{},
		interval: interval,
		ticker:   time.NewTicker(interval),
	}

	go newCache.reapLoop()

	return &newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.entries[key]; !exists {
		c.entries[key] = cacheEntry{
			createdAt: time.Now(),
			val:       val,
		}
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if val, exists := c.entries[key]; exists {
		return val.val, true
	}
	return nil, false
}

func (c *Cache) reapLoop() {
	for range c.ticker.C {
		c.mu.Lock()

		for key, entry := range c.entries {
			if time.Since(entry.createdAt) > c.interval {
				delete(c.entries, key)
			}
		}

		c.mu.Unlock()
	}
}
