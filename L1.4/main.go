package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const timeout = 3

func worker(ctx context.Context, wg *sync.WaitGroup, workerNum int, workerTime time.Duration) {
	defer wg.Done()
	log.Printf("worker %v started", workerNum)

	t := time.NewTicker(time.Second * workerTime)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			log.Printf("worker %v - working", workerNum)
		case <-ctx.Done():
			time.Sleep(time.Second * (workerTime + 1))
			log.Printf("worker %v stopped", workerNum)
			return
		}
	}
}

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())

	wg := &sync.WaitGroup{}

	for i := range 3 {
		wg.Add(1)
		workerNum := i + 1
		workerTime := time.Duration(i + 1)

		go worker(ctx, wg, workerNum, workerTime)

	}

	<-interrupt
	log.Println("waiting for workers...")
	cancel()

	doneChan := make(chan struct{})

	go func() {
		wg.Wait()
		close(doneChan)
	}()

	select {
	case <-time.After(time.Second * timeout):
		log.Println("shutdown by timeout")
	case <-doneChan:
		log.Printf("shutdown gracefully")
	}

}
