package main

import "sync"

type mutexCounter struct {
	val int64
	mu  sync.Mutex
}

func (c *mutexCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val++
}

func (c *mutexCounter) Value() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.val
}
