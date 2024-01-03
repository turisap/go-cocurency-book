package main

import (
	"fmt"
	"sync"
)

func errorHandleWaitGroup() {
	var e1 error
	var e2 error
	var wg sync.WaitGroup

	wg.Add(2)

	go func() {
		id := 1
		defer wg.Done()

		if res, err := computeResult(id); err != nil {
			fmt.Println("Got error in", id)
			e1 = err
		} else {
			fmt.Printf("RES in %d: %s", id, res)
		}
	}()

	go func() {
		id := 2
		defer wg.Done()

		if res, err := computeResult(id); err != nil {
			fmt.Println("Got error in", id)
			e1 = err
		} else {
			fmt.Printf("RES in %d: %s", id, res)
		}
	}()

	wg.Wait()
	if e1 != nil {
		fmt.Println("ERR1", e1)
	}
	if e2 != nil {
		fmt.Println("ERR2", e2)
	}

}
