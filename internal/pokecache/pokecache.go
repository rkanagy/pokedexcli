package pokecache

import (
	"fmt"
	"sync"
	"time"
)

// Cache represents a cache as a map with a Mutex
type Cache struct {
	cache map[string]CacheEntry
	mu    *sync.Mutex
}

// CacheEntry is what is stored as an entry in the cache
type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

// NewCache creates a new concurrency-safe cache
func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		cache: make(map[string]CacheEntry),
		mu:    &sync.Mutex{},
	}

	go newCache.reapLoop(interval)

	return newCache
}

// Add adds an entry into the cache
func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cache[key] = CacheEntry{
		val:       val,
		createdAt: time.Now(),
	}
}

// Get returns an entry from the cache, if found
func (c *Cache) Get(key string) (val []byte, found bool) {
	c.mu.Lock()

	entry, found := c.cache[key]

	c.mu.Unlock()

	var entryVal []byte
	if found {
		entryVal = entry.val
	}
	return entryVal, found
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer fmt.Println("ticker is being stopped")
	defer ticker.Stop()

	for {
		select {
		case t := <-ticker.C:
			c.releaseCacheEntries(t, interval)
		}
	}
}

func (c *Cache) releaseCacheEntries(currentTime time.Time, interval time.Duration) {
	c.mu.Lock()

	cutoffTime := currentTime.Add(-interval)
	for key, entry := range c.cache {
		if entry.createdAt.Before(cutoffTime) {
			delete(c.cache, key)
		}
	}

	c.mu.Unlock()
}
