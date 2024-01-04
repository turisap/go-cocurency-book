package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

func counter() {
	var done = make(chan struct{})
	var count int64
	nWorkers := 5

	for i := 1; i < nWorkers; i++ {
		go func(id int) {
			ticker := time.NewTicker(time.Millisecond * time.Duration(rand.Intn(1000)))
			defer ticker.Stop()

			for t := range ticker.C {
				_ = t
				atomic.AddInt64(&count, int64(id))
			}
		}(i)
	}

	go reader(&count, done)

	<-done
}
func reader(count *int64, done chan<- struct{}) {
	ticker := time.NewTicker(time.Millisecond * 10)
	defer ticker.Stop()

outer:
	for t := range ticker.C {
		val := atomic.LoadInt64(count)
		fmt.Printf("Got value %d at  %s\n", val, t.Format(time.Kitchen))

		if val > 10 {
			break outer
		}
	}

	fmt.Println("Sending done")
	close(done)
}
