package main

import (
	"fmt"
	"time"
)

func main() {
	//t1()
	t2()
}

func t1() {
	timer := time.NewTimer(time.Second)
	timeout := make(chan struct{})

	defer timer.Stop()

	go func() {
		<-timer.C
		close(timeout)
	}()

outer:
	for {
		select {
		case <-timeout:
			fmt.Println("Timeout")
			break outer
		default:
			fmt.Println("Observing sunset")
		}
	}

	fmt.Println("Finish execution")
}

func t2() {
	timeout := make(chan struct{})

	time.AfterFunc(900*time.Millisecond, func() {
		close(timeout)
		fmt.Println("Closing timeout")
	})

outer:
	for {
		select {
		case <-timeout:
			fmt.Println("Timeout")
			break outer
		default:
			time.Sleep(time.Millisecond * 80)
			fmt.Println("Observing sunset")
		}
	}

	fmt.Println("Finish execution")
}
