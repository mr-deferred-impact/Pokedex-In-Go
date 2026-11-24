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
	request map[string]cacheEntry
	mu      *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		request: make(map[string]cacheEntry),
		mu:      &sync.Mutex{},
	}

	go newCache.reapLoop(interval)

	return newCache
}

func (c *Cache) reap(maxEntryAge time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	cutoff := time.Now().Add(-maxEntryAge)

	for entry := range c.request {
		currentEntry := c.request[entry]

		if currentEntry.createdAt.Before(cutoff) {
			delete(c.request, entry)
		}
	}
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		c.reap(interval)
	}
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.request[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	returnedCacheEntry, status := c.request[key]

	return returnedCacheEntry.val, status
}
