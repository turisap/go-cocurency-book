package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//	func main() {
//		forks := [5]sync.Mutex{}
//
//		go philosopherDeadlocking(0, &forks[0], &forks[4])
//		go philosopherDeadlocking(1, &forks[0], &forks[1])
//		go philosopherDeadlocking(2, &forks[1], &forks[2])
//		go philosopherDeadlocking(3, &forks[2], &forks[3])
//		go philosopherDeadlocking(4, &forks[3], &forks[4])
//
//		select {}
//	}
func main() {
	forks := [5]chan bool{}

	for i, _ := range forks {
		forks[i] = make(chan bool, 1)
		forks[i] <- true
	}

	go philosopherChannelled(0, forks[0], forks[4])
	go philosopherChannelled(1, forks[0], forks[1])
	go philosopherChannelled(2, forks[1], forks[2])
	go philosopherChannelled(3, forks[2], forks[3])
	go philosopherChannelled(4, forks[3], forks[4])

	select {}
}

func philosopherDeadlocking(i int, lFork, rFork *sync.Mutex) {
	for {
		fmt.Printf("Philosopher %d is thinking\n", i)
		//time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))

		lFork.Lock()
		rFork.Lock()
		fmt.Printf("Philosopher %d is eating\n", i)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		lFork.Unlock()
		rFork.Unlock()
	}
}

func philosopherChannelled(i int, lFork, rFork chan bool) {
	for {
		fmt.Printf("Philosopher %d is thinking\n", i)
		select {
		case <-lFork:
			select {
			case <-rFork:
				fmt.Printf("Philosopher %d is eating\n", i)
				time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
				rFork <- true
			default:
				rFork <- true
			}
		}
		lFork <- true
	}
}
