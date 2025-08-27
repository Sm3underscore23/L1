package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"
)

var (
	workersNum = flag.Int("workers-num", 1, "Number of workers in worker pool")
)

func main() {
	flag.Parse()
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	dataChan := make(chan string, 1)
	wp := NewWP(*workersNum)
	wp.Create()

loop:
	for {
		select {
		case <-interrupt:
			break loop
		default:
			select {
			case dataChan <- "some info":
			case data := <-dataChan:
				wp.Handle(data)
			}
		}
	}

	wp.Whait()
}
