package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func fanOut(input <-chan int, numWorkers int) []<-chan int {
	channels := make([]<-chan int, numWorkers)
	for i := 0; i < numWorkers; i++ {
		channels[i] = worker(input)
	}
	return channels
}

func worker(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for n := range input {
			output <- process(n)
		}
	}()
	return output
}

func fanIn(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	multiplexedStream := make(chan int)

	multiplex := func(c <-chan int) {
		defer wg.Done()
		for i := range c {
			multiplexedStream <- i
		}
	}

	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	go func() {
		wg.Wait()
		close(multiplexedStream)
	}()

	return multiplexedStream
}

func process(n int) int {
	// Simulate some work
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	return n
}

func main() {

	var numOfTasks, numWorkers int
	fmt.Println("Enter the number of task")
	fmt.Scanln(&numOfTasks)
	fmt.Println("Enter the number of workers")
	fmt.Scanln(&numWorkers)
	input := make(chan int, numOfTasks)

	// Fan-out to n workers
	workers := fanOut(input, numWorkers)

	// Fan-in the results
	results := fanIn(workers...)

	// Send some input
	go func() {
		for i := 0; i < numOfTasks; i++ {
			input <- i
		}
		close(input)
	}()

	// Collect results
	for result := range results {
		fmt.Printf("Result: %d\n", result)
	}
}
