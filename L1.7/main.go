package main

import (
	"log"
	"sync"
)

type safeMap struct {
	data map[int]struct{}
	mu   sync.Mutex
}

func main() {

	m := safeMap{
		data: make(map[int]struct{}),
	}

	wg := sync.WaitGroup{}

	for i := range 3 {
		for range 2 {
			wg.Add(1)
			i := i
			go func() {
				defer wg.Done()
				m.mu.Lock()
				defer m.mu.Unlock()
				_, ok := m.data[i]
				if !ok {
					m.data[i] = struct{}{}
					log.Printf("number %v add", i)
					return
				}
				log.Printf("number %v already exist", i)
			}()
		}
	}

	wg.Wait()
}
