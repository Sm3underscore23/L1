package main

import (
	"fmt"
	"sync"
)

type counter interface {
	Inc()
	Value() int64
}

func withAtomic() {
	c := counter(&atomicCounter{})
	wg := sync.WaitGroup{}

	for range 10 {

		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()

	fmt.Println("with atomic", c.Value())
}

func withMutex() {
	c := counter(&mutexCounter{})

	wg := sync.WaitGroup{}

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c.Inc()
		}()
	}

	wg.Wait()

	fmt.Println("with mutex", c.Value())
}

func main() {
	withAtomic()
	withMutex()
}
