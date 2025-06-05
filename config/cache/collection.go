package cache

import (
	"fmt"
	"reflect"
	"sync"
)

type Collection struct {
	data  map[string]interface{}
	count uint64
	mut   *sync.RWMutex
}

func newCollection() *Collection {
	return &Collection{
		data:  make(map[string]interface{}),
		count: 0,
		mut:   &sync.RWMutex{},
	}
}

func (c *Collection) Add(key string, value interface{}) error {
	c.mut.Lock()
	defer c.mut.Unlock()

	if c.data[key] != nil {
		return fmt.Errorf("key [%s] already exists", key)
	}

	c.data[key] = value
	c.count++
	return nil
}

func (c *Collection) AddOrReplace(key string, value interface{}) {
	c.mut.Lock()
	defer c.mut.Unlock()

	c.data[key] = value
	c.count++
}

func (c *Collection) Get(key string, model interface{}) error {
	c.mut.RLock()
	defer c.mut.RUnlock()

	if value, ok := c.data[key]; !ok {
		return fmt.Errorf("key [%s] not found", key)
	} else {
		reflect.ValueOf(model).Elem().Set(reflect.ValueOf(value))
		return nil
	}
}

func (c *Collection) GetOrDefault(key string, model interface{}) interface{} {
	c.mut.RLock()
	defer c.mut.RUnlock()

	if value, ok := c.data[key]; !ok {
		return nil
	} else {
		reflect.ValueOf(model).Elem().Set(reflect.ValueOf(value))
		return &value
	}
}

func (c *Collection) Count() uint64 {
	return c.count
}

func (c *Collection) Remove(key string) {
	c.mut.Lock()
	defer c.mut.Unlock()

	if _, ok := c.data[key]; !ok {
		return
	}

	delete(c.data, key)
	c.count--
}

func (c *Collection) Clear() {
	c.mut.Lock()
	defer c.mut.Unlock()

	c.data = make(map[string]interface{})
	c.count = 0
}
