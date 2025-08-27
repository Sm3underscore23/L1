package main

import (
	"sync"
	"sync/atomic"
)

func withAtomic(data []int) int {
	var counter atomic.Int64
	wg := sync.WaitGroup{}

	for _, num := range data {
		wg.Add(1)
		num := num
		go func() {
			defer wg.Done()
			counter.Add(int64(num * num))
		}()
	}

	wg.Wait()

	return int(counter.Load())
}
