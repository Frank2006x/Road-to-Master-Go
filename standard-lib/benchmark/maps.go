package main

import (
	"sync"
)

// MapWithMutex is a thread-safe map using mutex
type MapWithMutex struct {
	mu    sync.RWMutex
	items map[string]int
}

func NewMapWithMutex() *MapWithMutex {
	return &MapWithMutex{
		items: make(map[string]int),
	}
}

func (m *MapWithMutex) Set(key string, value int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.items[key] = value
}

func (m *MapWithMutex) Get(key string) (int, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	val, ok := m.items[key]
	return val, ok
}

func (m *MapWithMutex) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.items, key)
}

// SyncMapImpl wraps sync.Map
type SyncMapImpl struct {
	items sync.Map
}

func NewSyncMap() *SyncMapImpl {
	return &SyncMapImpl{}
}

func (m *SyncMapImpl) Set(key string, value int) {
	m.items.Store(key, value)
}

func (m *SyncMapImpl) Get(key string) (int, bool) {
	val, ok := m.items.Load(key)
	if !ok {
		return 0, false
	}
	return val.(int), true
}

func (m *SyncMapImpl) Delete(key string) {
	m.items.Delete(key)
}

// RawMap is a non-concurrent map (baseline)
type RawMap struct {
	items map[string]int
}

func NewRawMap() *RawMap {
	return &RawMap{
		items: make(map[string]int),
	}
}

func (m *RawMap) Set(key string, value int) {
	m.items[key] = value
}

func (m *RawMap) Get(key string) (int, bool) {
	val, ok := m.items[key]
	return val, ok
}

func (m *RawMap) Delete(key string) {
	delete(m.items, key)
}
