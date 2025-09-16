package main

import (
	"fmt"
	"sync"
)

func square(data []int) {
	wg := sync.WaitGroup{}

	for _, num := range data {
		wg.Add(1)
		num := num
		go func() {
			defer wg.Done()
			fmt.Println(num * num)
		}()
	}

	wg.Wait()
}

func main() {
	data := []int{2, 4, 6, 8, 1}
	square(data)
}
