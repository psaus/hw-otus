package hw04lrucache

import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
	sync.Mutex
}

type cacheItem struct {
	key   Key
	value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.Lock()
	defer c.Unlock()

	listItem, ok := c.items[key]
	ci := &cacheItem{key, value}

	if !ok {
		c.items[key] = c.queue.PushFront(ci)
		c.Clear()
	} else {
		listItem.Value = ci
		c.queue.MoveToFront(listItem)
	}

	return ok
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.Lock()
	defer c.Unlock()

	v, ok := c.items[key]

	if v != nil && ok {
		c.queue.MoveToFront(v)
		switch ci := v.Value.(type) {
		case *cacheItem:
			return ci.value, true
		default:
			return v.Value, true
		}
	}

	return nil, false
}

func (c *lruCache) Clear() {
	if c.queue.Len() > c.capacity {
		itemForRemove := c.queue.Back()
		c.queue.Remove(itemForRemove)
		cacheForRemove := itemForRemove.Value.(*cacheItem)
		delete(c.items, cacheForRemove.key)
	}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
