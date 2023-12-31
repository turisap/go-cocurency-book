package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	sharedCounter := atomic.Int64{}
	sharedCounter.Store(0)
	c1 := make(chan bool)
	c2 := make(chan bool)

	go incrementor(&sharedCounter, c1)

	go decrementor(&sharedCounter, c2)

	<-c1
	<-c2

	fmt.Println(sharedCounter.Load())
}

func incrementor(c *atomic.Int64, done chan bool) {
	for i := 0; i < 10000; i++ {
		c.Add(1)
	}
	done <- true
}

func decrementor(c *atomic.Int64, done chan bool) {
	for i := 0; i < 10000; i++ {
		c.Add(-1)
	}
	done <- true
}
