package main

import (
	"fmt"
	"sync"
)

var x = 0
var ch = make(chan int)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for {
			x++
			ch <- 1
		}
		wg.Done()
	}()
	go func() {
		for range ch {
			fmt.Println(x)
		}
		wg.Done()
	}()

	wg.Wait()

}
