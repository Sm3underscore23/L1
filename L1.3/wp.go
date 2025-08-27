package main

import (
	"log"
	"time"
)

type worker struct {
	name     int
	jobsDone int
}

type workerPool struct {
	workers    chan worker
	handleTime time.Duration
}

func NewWP(workersNum int, handleTime time.Duration) *workerPool {
	return &workerPool{
		workers:    make(chan worker, workersNum),
		handleTime: handleTime,
	}
}

func (wp *workerPool) Create() {
	for i := range cap(wp.workers) {
		wp.workers <- worker{name: i}
	}
}

func (wp *workerPool) Handle(data string) {
	worker := <-wp.workers
	go func() {
		log.Printf("worker %v: processing...", worker.name)
		time.Sleep(wp.handleTime)
		log.Printf("worker %v: done", worker.name)
		worker.jobsDone++
		wp.workers <- worker
	}()
}

func (wp *workerPool) Whait() {
	for range cap(wp.workers) {
		worker := <-wp.workers
		log.Printf("worker %v: jobs colpleted: %v", worker.name, worker.jobsDone)
	}
}
