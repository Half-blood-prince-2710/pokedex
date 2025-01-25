package pokecache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val []byte
}

type Cache struct {
	mp map[string]cacheEntry
	mu sync.Mutex
}

func NewCache(interval time.Duration) *Cache{
	cache := &Cache{
		mp: make(map[string]cacheEntry),
	}
	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte ){
	c.mu.Lock()
	defer c.mu.Unlock()
	c.mp[key] = cacheEntry{val:val,createdAt: time.Now()}
} 

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _,exists:=c.mp[key]; exists {
		return c.mp[key].val, true
	}else {
		return nil, false
	}

}

func (c *Cache) reapLoop(interval time.Duration) {

	 for {
			time.Sleep(interval)
			c.mu.Lock()
			if len(c.mp)>0 {
				for key, _ := range c.mp {
			if time.Since(c.mp[key].createdAt)> interval {
				delete(c.mp,key)
			}
		}
			}
		
		c.mu.Unlock()
	}
}
