package main

import "sync/atomic"

type atomicCounter struct {
	val int64
}

func (c *atomicCounter) Inc() {
	atomic.AddInt64(&c.val, 1)
}

func (c *atomicCounter) Value() int64 {
	return atomic.LoadInt64(&c.val)
}
