package main

import (
	"log"
	"sync"
	"time"
)

const (
	timeout = 1
)

func main() {
	t := time.NewTimer(time.Second * timeout)

	dataChan := make(chan struct{})

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for {
			select {
			case dataChan <- struct{}{}:
				log.Println("write")
				time.Sleep(time.Millisecond * 200)
			case <-t.C:
				log.Println("writer stopped")
				close(dataChan)
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		for range dataChan {
			log.Println("read")
		}
		log.Println("reader stopped")
	}()

	wg.Wait()
}
