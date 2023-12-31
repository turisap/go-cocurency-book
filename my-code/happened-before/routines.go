package main

import (
	"fmt"
	"sync"
)

func main() {
	c := 0

	for c < 1000 {
		printOneOrZero()
	}
}

func printOneOrZero() {
	var wg sync.WaitGroup
	var x int
	wg.Add(1)
	go func() {
		x = 1
		wg.Done()
	}()
	if x == 0 {
		fmt.Println(x)
	} else {
		fmt.Println("OOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOOO", x)
	}
	wg.Wait()
}

func x10() {
	var x int
	ch := make(chan int)
	go func() {
		ch <- 0
		// always prints 1
		// because x is written before
		// channel read
		fmt.Println(x)
	}()
	x = 1
	<-ch
	select {}
}
