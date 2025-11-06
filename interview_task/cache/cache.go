package cache

import "sync"

type Cache struct {
	Data map[string]any
	Mtx  sync.RWMutex
}

func (c *Cache) Store(key string, value any) {
	c.Mtx.Lock()
	defer c.Mtx.Unlock()
	c.Data[key] = value
}

func (c *Cache) Load(key string) (any, bool) {
	c.Mtx.RLock()
	defer c.Mtx.RUnlock()
	v, ok := c.Data[key]
	return v, ok
}
