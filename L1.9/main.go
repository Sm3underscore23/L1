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

func readFromArray(
	ctx context.Context, cancel context.CancelFunc, interrupt chan os.Signal,
	data [10]int, producerCh, consumerCh chan int,
) {
	go func() {
		for _, x := range data {
			select {
			case <-ctx.Done():
				close(producerCh)
				return
			case producerCh <- x:
				time.Sleep(time.Millisecond * 300)
			}
		}
		close(producerCh)
	}()

	go func() {
		for x := range producerCh {
			consumerCh <- x * 2
		}
		close(consumerCh)
	}()

	go func() {
		<-interrupt
		cancel()
	}()

	for result := range consumerCh {
		fmt.Println(result)
	}

	fmt.Println("done")
}

func generateFromArray(
	ctx context.Context, cancel context.CancelFunc, interrupt chan os.Signal,
	data [10]int, producerCh, consumerCh chan int,
) {
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

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	data := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	producerCh := make(chan int)
	consumerCh := make(chan int)

	// readFromArray(ctx, cancel, interrupt, data, producerCh, consumerCh)
	generateFromArray(ctx, cancel, interrupt, data, producerCh, consumerCh)
}
