package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entryMap map[string]cacheEntry
	mu       sync.RWMutex
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(duration time.Duration) *Cache {
	cache := Cache{}
	cache.entryMap = make(map[string]cacheEntry)
	go cache.reapLoop(duration)
	return &cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	cache.entryMap[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mu.RLock()
	defer cache.mu.RUnlock()
	val, ok := cache.entryMap[key]
	if ok {
		return val.val, true
	} else {
		return []byte{}, false
	}
}

func (cache *Cache) reapLoop(duration time.Duration) {
	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	for range ticker.C {
		cache.reap(duration)
	}
}

func (cache *Cache) reap(interval time.Duration) {
	cache.mu.Lock()
	defer cache.mu.Unlock()
	for key, val := range cache.entryMap {
		if interval.Milliseconds() < time.Since(val.createdAt).Milliseconds() {
			delete(cache.entryMap, key)
		}
	}
}
