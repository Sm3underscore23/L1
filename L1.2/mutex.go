package main

import "sync"

func withMutex(data []int) int {
	counter := 0

	m := sync.Mutex{}
	wg := sync.WaitGroup{}

	for _, num := range data {
		wg.Add(1)
		num := num
		go func() {
			defer wg.Done()
			m.Lock()
			defer m.Unlock()
			counter += num * num
		}()
	}

	wg.Wait()

	return counter
}
