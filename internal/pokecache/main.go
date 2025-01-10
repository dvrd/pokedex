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
	data map[string]cacheEntry
	lock sync.RWMutex
}

func NewCache(interval time.Duration) *Cache {
	newCache := Cache{
		data: map[string]cacheEntry{},
	}

	go newCache.reapLoop(interval)

	return &newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.lock.Lock()
	defer c.lock.Unlock()

	c.data[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.lock.RLock()
	defer c.lock.RUnlock()

	entry, ok := c.data[key]

	return entry.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	then := time.Now()
	time.Sleep(interval)

	c.lock.RLock()
	defer c.lock.RUnlock()

	for key, val := range c.data {
		if then.Sub(val.createdAt) > 0 {
			delete(c.data, key)
		}
	}
}
