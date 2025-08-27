package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	dataChan := make(chan string, 1)
	wp := NewWP(4)
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
