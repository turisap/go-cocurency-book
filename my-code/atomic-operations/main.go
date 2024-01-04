package main

import (
	"fmt"
	"sync/atomic"
)

func main() {
	//i := 0
	//for i < 1000 {
	//	//memoryRace()
	//	//noMemoryRace()
	//	dataRace()
	//	i++
	//}

	//counter()
	//orchestrator()
	processUpdate()
}

func badCriticalExample() {
	var locked atomic.Bool

	// it is false here
	if !locked.Load() {
		// but here it can be true already
		// cause another goroutine set it
		locked.Store(true)
		defer locked.Store(false)
		// do logic
	}

	// SOLUTION
	if !locked.CompareAndSwap(false, true) {
		defer locked.Store(false)
		//do logic
	}
}

func noMemoryRace() {
	var done atomic.Bool
	var str atomic.Value

	go func() {
		str.Store("Hey")
		done.Store(true)
	}()

	for !done.Load() {

	}

	fmt.Println(done.Load())
}

func memoryRace() {
	var str string
	var done bool
	go func() {
		str = "Done!"
		done = true
	}()
	for !done {
	}
	fmt.Println(str)
}

func dataRace() {
	var done atomic.Bool
	var a int
	go func() {
		a = 1
		done.Store(true)
	}()
	// done can be false here
	if done.Load() {
		// but can  be already true here
		fmt.Println(a)
	}
}
