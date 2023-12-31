package main

import (
	"fmt"
	"sync"
)

func main() {
	var i int
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		i = 1
		wg.Done()
	}()
	wg.Wait()
	// If we are here, wg.Done is called, so i=1
	fmt.Println(i)
}
