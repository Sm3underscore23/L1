package main

import (
	"sync"
)

func withChan(data []int) int {
	nums := make(chan int, len(data))
	wg := sync.WaitGroup{}

	for _, num := range data {
		wg.Add(1)
		num := num
		go func() {
			defer wg.Done()
			nums <- num * num
		}()
	}

	go func() {
		wg.Wait()
		close(nums)
	}()

	counter := 0

	for num := range nums {
		counter += num
	}

	return counter
}
