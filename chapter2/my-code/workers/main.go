package main

import (
	"fmt"
	"time"
)

const nWorkers = 5 // Replace this with the desired number of workers
const nJobs = 30   // Replace this with the total number of jobs

func main() {
	cW := make(chan int)
	cR := make(chan int)
	done := make(chan bool)

	for i := 0; i < nWorkers; i++ {
		go worker(i, cW, cR)
	}

	go collectResults(cR, done)

	for i := 0; i < nJobs; i++ {
		cW <- i
	}

	close(cW)

	fmt.Println("All done", <-done)
}

func worker(i int, cw, cr chan int) {
	for job := range cw {
		fmt.Printf("Worker %d got a job %d\n", i, job)
		time.Sleep(time.Millisecond * 100)
		cr <- job
		fmt.Printf("Worker %d send the result %d back\n", i, job)
	}
}

func collectResults(cR chan int, done chan bool) {
	c := 0
	for res := range cR {
		fmt.Println("collected result ", res)
		c++

		if c >= nJobs {
			break
		}
	}

	done <- true
	close(cR)
	close(done)
}
