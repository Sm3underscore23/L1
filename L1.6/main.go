package main

import (
	"context"
	"log"
	"runtime"
	"sync"
	"time"
)

const (
	byReturn = iota + 1
	byGoexit
	byChan
	byContext
	byTimer
	byGlobalFlag
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*byContext)
	defer cancel()

	exitChan := make(chan struct{})

	wg := sync.WaitGroup{}

	wg.Add(7)

	go func() {
		defer wg.Done()
		time.Sleep(time.Second * byChan)
		close(exitChan)
	}()

	go func() {
		defer wg.Done()
		time.Sleep(time.Second * byReturn)
		if 2+2 == 4 {
			log.Println("by return")
			return
		}
		log.Println("no data")
	}()

	go func() {
		time.Sleep(time.Second * byGoexit)
		log.Println("by runtime.Goexit()")
		wg.Done()
		runtime.Goexit()
		log.Println("no data")

	}()

	go func() {
		defer wg.Done()
		<-exitChan
		log.Println("by chan")
	}()

	go func() {
		defer wg.Done()
		<-ctx.Done()
		log.Println("by contex timeout")
	}()

	go func() {
		defer wg.Done()
		<-time.After(time.Second * byTimer)
		log.Println("by time.After()")
	}()

	var f bool

	go func() {
		defer wg.Done()
		time.Sleep(time.Second * byGlobalFlag)
		if f {
			log.Println("by global flag")
			return
		}
		log.Println("no data")
	}()

	f = true

	wg.Wait()
}
