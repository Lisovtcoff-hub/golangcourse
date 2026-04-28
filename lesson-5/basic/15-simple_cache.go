package main

import "sync"

type Cache struct {
	mx sync.RWMutex
	m  map[string]string
}

func NewCache() *Cache {
	return &Cache{
		m: make(map[string]string),
	}
}

func (c *Cache) Add(key, value string) {
	c.mx.Lock()
	defer c.mx.Unlock()

	c.m[key] = value
}

func (c *Cache) Get(key string) (string, bool) {
	c.mx.RLock()
	defer c.mx.RUnlock()

	value, ok := c.m[key]

	return value, ok
}
