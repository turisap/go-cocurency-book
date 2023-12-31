package main

import (
	"sync"
	"time"
)

type Cache struct {
	values sync.Map
}
type cachedValue struct {
	sync.Once
	value *Data
}

type Data struct {
	bytes []byte
}

func (c *Cache) Get(id string) *Data {
	// Get the cached value, or store an empty value
	v, _ := c.values.LoadOrStore(id, &cachedValue{})
	cv := v.(*cachedValue)
	// If not initialized, initialize here
	cv.Do(func() {
		time.Sleep(time.Second)
		cv.value = &Data{bytes: []byte{12, 3, 4}}
	})
	return cv.value
}
