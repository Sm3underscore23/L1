package main

import (
	"log"
	"time"
)

type worker struct {
	name     int
	jobsDone int
}

type wp struct {
	workers chan worker
}

func NewWP(workersNum int) *wp {
	return &wp{
		workers: make(chan worker, workersNum),
	}
}

func (wp *wp) Create() {
	for i := range cap(wp.workers) {
		wp.workers <- worker{name: i}
	}
}

func (wp *wp) Handle(data string) {
	worker := <-wp.workers
	go func() {
		log.Printf("worker %v: processing...", worker.name)
		time.Sleep(5 * time.Second)
		log.Printf("worker %v: done", worker.name)
		worker.jobsDone++
		wp.workers <- worker
	}()
}

func (wp *wp) Whait() {
	for range cap(wp.workers) {
		worker := <-wp.workers
		log.Printf("worker %v: jobs done: %v", worker.name, worker.jobsDone)
	}
}
