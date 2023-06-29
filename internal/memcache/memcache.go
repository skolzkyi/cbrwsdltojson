package memcache

import (
	"sync"
)

type MemCache struct {
	mu    sync.RWMutex
	cache map[string]interface{}
}

func New() *MemCache {
	return &MemCache{}
}

func (mc *MemCache) Init() {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	mc.cache = make(map[string]interface{})
}

func (mc *MemCache) AddOrUpdatePayloadInCache(tag string, payload interface{}) bool {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	_, ok := mc.cache[tag]
	mc.cache[tag] = payload
	// true is update
	return ok
}

func (mc *MemCache) RemovePayloadInCache(tag string) {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	delete(mc.cache, tag)
}

func (mc *MemCache) GetPayloadInCache(tag string) (interface{}, bool) {
	mc.mu.RLock()
	defer mc.mu.RUnlock()
	_, ok := mc.cache[tag]
	if !ok {
		return nil, ok
	}
	return mc.cache[tag], ok
}
