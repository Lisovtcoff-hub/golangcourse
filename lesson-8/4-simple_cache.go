package main

import (
	"errors"
	"fmt"
	"sync"
)

var ErrNotFound = errors.New("not found")

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

func (c *Cache) Get(key string) (string, error) {
	c.mx.RLock()
	defer c.mx.RUnlock()

	value, ok := c.m[key]
	if !ok {
		return "", ErrNotFound
	}

	return value, nil
}

func main() {
	cache := NewCache()

	cache.Add("key", "value")

	value, err := cache.Get("key1")
	if err != nil {
		fmt.Println("error:", err)

		return
	}

	fmt.Println("success:", value)
}
