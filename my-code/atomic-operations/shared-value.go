package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Person struct {
	id int
}

func computeNewCopy(in Person) Person {
	return Person{id: in.id + 1}
}

var sharedValue atomic.Pointer[Person]

func processUpdate() {
	var wg sync.WaitGroup
	sharedValue.Store(&Person{id: 1})

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go updateSharedValue(&wg)
	}

	wg.Wait()
	fmt.Println(sharedValue.Load())
}

func updateSharedValue(wg *sync.WaitGroup) {
	myCopy := sharedValue.Load()
	newCopy := computeNewCopy(*myCopy)

	if sharedValue.CompareAndSwap(myCopy, &newCopy) {
		fmt.Println("Set value successful")
	} else {
		fmt.Println("Another goroutine modified the value")
	}
	wg.Done()
}
