package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var s atomic.Value
	done := make(chan bool)

	go store(&s)
	go check(&s, done)

	<-done
}

func store(s *atomic.Value) {
	time.Sleep(time.Second * 2)

	s.Store(5)
}

func check(s *atomic.Value, done chan<- bool) {
	ticker := time.NewTicker(time.Millisecond * 500)

outer:
	for tick := range ticker.C {
		fmt.Println("Tick", tick.Format(time.Kitchen))

		if v := s.Load(); v == 5 {
			fmt.Println("found 5")
			break outer
		}
	}

	done <- true
}
