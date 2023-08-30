package memcache

import (
	"fmt"
	"sync"
	"time"
)

type CacheInfo struct {
	Payload     interface{}
	InfoDTStamp time.Time
}

type MemCache struct {
	mu    sync.RWMutex
	cache map[string]CacheInfo
}

func New() *MemCache {
	return &MemCache{}
}

func (mc *MemCache) Init() {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	mc.cache = make(map[string]CacheInfo)
}

func (mc *MemCache) AddOrUpdatePayloadInCache(tag string, payload interface{}) bool {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	tempEl, ok := mc.cache[tag]
	tempEl.Payload = payload
	tempEl.InfoDTStamp = time.Now()
	mc.cache[tag] = tempEl
	// true is update
	return ok
}

func (mc *MemCache) RemovePayloadInCache(tag string) {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	delete(mc.cache, tag)
}

func (mc *MemCache) GetCacheDataInCache(tag string) (CacheInfo, bool) {
	mc.mu.RLock()
	defer mc.mu.RUnlock()
	_, ok := mc.cache[tag]
	if !ok {
		return CacheInfo{}, ok
	}
	return mc.cache[tag], ok
}

func (mc *MemCache) RemoveAllPayloadInCacheByTimeStamp(controlTime time.Time) {
	mc.mu.Lock()
	defer mc.mu.Unlock()
	for key, data := range mc.cache {
		if data.InfoDTStamp.Before(controlTime) {
			delete(mc.cache, key)
		}
	}
}

func (mc *MemCache) PrintAllCacheKeys() {
	mc.mu.RLock()
	defer mc.mu.RUnlock()
	for key := range mc.cache {
		fmt.Println(key)
	}
}
