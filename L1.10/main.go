package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"math"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
)

const workersNum = 10

type result struct {
	checker   map[float64]struct{}
	tempsData map[string][]float64
	mu        sync.Mutex
}

func handle(rowTemp string, result *result) error {
	temp, err := strconv.ParseFloat(rowTemp, 64)
	if err != nil {
		return err
	}

	var group string
	switch {
	case temp >= 0 && temp < 10:
		group = "0"
	case temp < 0 && temp > -10:
		group = "-0"
	case temp >= 10:
		group = fmt.Sprintf("%d", int(math.Floor(temp))/10*10)
	case temp <= -10:
		fmt.Println(int(temp))
		group = fmt.Sprintf("%d", int(temp)/10*10)
	}

	result.mu.Lock()
	defer result.mu.Unlock()
	if _, ok := result.checker[temp]; ok {
		fmt.Printf("%v alredy exist\n", temp)
		return nil
	}

	result.checker[temp] = struct{}{}
	result.tempsData[group] = append(result.tempsData[group], temp)
	return nil
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	r := bufio.NewReader(os.Stdin)

	result := result{
		checker:   map[float64]struct{}{},
		tempsData: map[string][]float64{},
	}

	data, _, err := r.ReadLine()
	if err != nil {
		log.Fatal(err)
	}

	rowTemps := strings.Split(string(data), ", ")

	wp := make(chan struct{}, workersNum)
	for range workersNum {
		wp <- struct{}{}
	}

	go func() {
		<-interrupt
		cancel()
	}()

l:
	for _, rowTemp := range rowTemps {
		select {
		case <-ctx.Done():
			break l
		default:
			<-wp
			go func() {
				if err := handle(rowTemp, &result); err != nil {
					log.Fatal(err)
				}
				wp <- struct{}{}
			}()
		}
	}

	for range cap(wp) {
		<-wp
	}

	fmt.Println(result.tempsData)
}
