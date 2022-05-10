package main

import (
	"errors"
	"sync"
)

type iCache interface {
	Item
}

type Cache[T iCache] struct {
	data map[string]T
	lock *sync.Mutex
}

var singletonInsts = make(map[string]*Cache[Item])
var lock = &sync.Mutex{}

func GetInstance[T Item](component string) *Cache[Item] {
	if _, ok := singletonInsts[component]; !ok {
		lock.Lock()
		defer lock.Unlock()
		{
			if _, ok := singletonInsts[component]; !ok {
				singletonInsts[component] = &Cache[Item]{data: make(map[string]Item), lock: &sync.Mutex{}}
			}
		}

	}
	return singletonInsts[component]
}

func (c *Cache[T]) GetValue(key string) (T, error) {
	val, ok := c.data[key]
	if ok {
		return val, nil
	}
	return val, errors.New("not found")

}

func (c *Cache[T]) SetValue(key string, val T) {
	c.lock.Lock()
	defer c.lock.Unlock()
	c.data[key] = val
}

func (c *Cache[T]) DeleteValue(key string) {
	c.lock.Lock()
	defer c.lock.Unlock()
	delete(c.data, key)
}
