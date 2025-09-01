package main

import (
	"fmt"
	"sync"
	"time"
)

func customTimeSleep(d time.Duration) {
	t := time.NewTimer(d)
	<-t.C
}

func main() {
	wg := sync.WaitGroup{}

	for i := range 3 {
		wg.Add(1)
		i := i + 1
		go func() {
			defer wg.Done()
			now := time.Now()
			fmt.Printf("%v worker started\n", i)
			customTimeSleep(time.Second * time.Duration(i))
			fmt.Printf("%v worker done - time sleep: %v\n", i, time.Since(now))
		}()
	}
	wg.Wait()
}
