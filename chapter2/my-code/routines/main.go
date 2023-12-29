package main

import (
	"fmt"
	"time"
)

type Data int

// the receiving end can get the data
// but cannot send any. This is type-safety
func streamResults() <-chan Data {
	resultCh := make(chan Data)

	go func() {
		defer close(resultCh)
		results := []Data{1, 2, 3, 4, 5}
		for _, result := range results {
			resultCh <- result
		}
	}()

	return resultCh
}

func main() {
	ch := streamResults()

	for v := range ch {
		fmt.Println("Value received", v)
		time.Sleep(500 * time.Millisecond)
	}
}

//
//func f(s string) {
//	fmt.Printf("Goroutine %s\n", s)
//}
//func main() {
//	// prints out c c c because of closure
//	// s escapes to the heap
//	for _, s := range []string{"a", "b", "c"} {
//		go func() {
//			fmt.Printf("Goroutine %s\n", s)
//		}()
//	}
//	// prints out abc in any order
//	for _, s := range []string{"a", "b", "c"} {
//		go func(str string) {
//			fmt.Printf("Goroutine %s\n", str)
//		}(s)
//	}
//	time.Sleep(900000)
//}
