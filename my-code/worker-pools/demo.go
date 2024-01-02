package main

import (
	"fmt"
	"time"
)

type Product string

type Job struct {
	id       int
	duration int
	results  chan Product
}

func main() {
	jobs := make(chan Job)
	allResults := make(chan chan Product)

	// create worker pool
	for i := 0; i < 3; i++ {
		go func(i int, jobs chan Job) {
			for job := range jobs {
				fmt.Printf("[#%d]JOB: %v\n", job, i)

				for j := 0; j < job.duration; j++ {
					fmt.Println("Pushing job")
					job.results <- Product(fmt.Sprintf("Result from #%d worker, res %d", i, j))
				}
				close(job.results)
			}
			fmt.Println("Finish worker")
		}(i, jobs)
	}

	// producer
	go func(j chan Job) {
		defer close(allResults)
		ticker := time.NewTicker(time.Millisecond * 100)
		c := 0

		for tick := range ticker.C {
			fmt.Println("Ticking")
			res := make(chan Product)
			jobs <- Job{
				id:       int(tick.UnixMilli()),
				duration: c,
				results:  res,
			}
			allResults <- res

			if c > 20 {
				fmt.Println("finish producing")
				ticker.Stop()
				close(j)
				break
			}
			c++
		}
	}(jobs)

	for resCh := range allResults {
		for res := range resCh {
			fmt.Println("RES", res)
		}
	}

}

// workerMy

// producer
