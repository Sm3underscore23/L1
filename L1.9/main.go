package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	data := [5]int{1, 2, 3, 4, 5}

	producerCh := make(chan int)

	consumerCh := make(chan int)

	wg := sync.WaitGroup{}

	wg.Add(3)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				close(producerCh)
				return
			case producerCh <- data[rand.Intn(len(data))]:
				time.Sleep(time.Millisecond * 150)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for x := range producerCh {
			consumerCh <- x * 2
		}
		close(consumerCh)
	}()

	go func() {
		defer wg.Done()
		for result := range consumerCh {
			fmt.Println(result)
		}
	}()

	<-interrupt
	cancel()
	wg.Wait()
	fmt.Println("stopped")
}
