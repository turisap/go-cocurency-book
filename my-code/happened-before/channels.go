package main

import (
	"fmt"
	"sync"
)

type Work = []int

func main() {
	resources := make(chan struct{}, 5)
	jobs := make(chan Work)

	var wg = sync.WaitGroup{}

	for worker := 0; worker < 10; worker++ {
		wg.Add(1)
		go fn(&wg, resources, jobs)
	}

	wg.Wait()
}
func fn(waitGroup *sync.WaitGroup, resources chan struct{}, jobs chan Work) {
	for work := range <-jobs {
		// Do work
		// Acquire resource
		resources <- struct{}{}
		// Work with resource
		<-resources

		fmt.Println("Oh")

		_ = work
		waitGroup.Done()
	}
}
