package pokecache

import (
	"sync"
	"time"
)

type RestApiLocation struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Cache struct {
	cache map[string]CacheEntry
	mux   *sync.Mutex
}

type CacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]CacheEntry),
		mux:   &sync.Mutex{},
	}

	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = CacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}

}

func (c *Cache) Get(key string) ([]byte, bool) {

	c.mux.Lock()
	defer c.mux.Unlock()
	entry, ok := c.cache[key]

	return entry.val, ok

}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.cache {
		if now.Sub(v.createdAt) > last {
			delete(c.cache, k)
		}
	}
}
