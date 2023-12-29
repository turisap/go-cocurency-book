package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Work struct {
	value int
}

type Result struct {
	value int
}

func main() {
	workCh := make(chan Work)
	resultCh := make(chan Result)
	done := make(chan bool)

	workQueue := make([]Work, 100)
	for i := range workQueue {
		workQueue[i].value = rand.Int()
	}

	// Create 10 worker goroutines
	for i := 0; i < 10; i++ {
		go func() {
			for {
				// Get work from the work channel
				fmt.Println("got a job")
				work := <-workCh
				// Compute result
				result := Result{
					value: work.value * 2,
				}
				// Send the result via the result channel
				time.Sleep(time.Second)
				resultCh <- result
				fmt.Println("finished job")
			}
		}()
	}
	results := make([]Result, 0)
	go func() {
		// Collect all the results.
		for i := 0; i < len(workQueue); i++ {
			results = append(results, <-resultCh)
			fmt.Println("got a result")
		}
		// When all the results are collected, notify the done channel
		done <- true
	}()
	// Send all the work to the workers
	for _, work := range workQueue {
		workCh <- work
	}

	close(workCh)
	// Wait until everything is done
	<-done
	fmt.Println(results)
}
