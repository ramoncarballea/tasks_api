package cache

import (
	"fmt"
	"sync"
)

type memoryCache struct {
	Storage
	data map[string]*Collection
	mut  *sync.RWMutex
}

func NewMemory() Storage {
	return &memoryCache{
		data: make(map[string]*Collection),
		mut:  &sync.RWMutex{},
	}
}

func (m *memoryCache) Add(key string) {
	m.mut.Lock()
	defer m.mut.Unlock()

	m.data[key] = newCollection()
}

func (m *memoryCache) TryGet(key string) (*Collection, error) {
	m.mut.RLock()
	defer m.mut.RUnlock()

	val := m.data[key]
	if val == nil {
		return nil, fmt.Errorf("key [%s] not found", key)
	}
	return val, nil
}

func (m *memoryCache) Clear() {
	m.data = make(map[string]*Collection)
}

func (m *memoryCache) Remove(key string) {
	m.mut.Lock()
	defer m.mut.Unlock()

	delete(m.data, key)
}
