package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c2 := make(chan rune)
	c1 := make(chan rune)
	done := make(chan bool)
	c := 0

	go send(c1, 300, &c)
	go send(c2, 200, &c)
	go terminate(done, &c)

outer:
	for {
		select {
		case v := <-c1:
			fmt.Println("Got from #1", v)
		case v := <-c2:
			fmt.Println("Got from #2", v)
		case <-done:
			fmt.Println("Finish processing")
			break outer
		default:
			fmt.Println("Default")
			time.Sleep(time.Millisecond * 20)
		}

	}

	fmt.Println("All done", <-done)
}

func send(ch chan rune, d time.Duration, c *int) {

	source := rand.NewSource(time.Now().UnixNano())
	randGen := rand.New(source)
	randomRune := rune(randGen.Intn(0x10FFFF))

	time.Sleep(time.Millisecond * d)
	*c = *c + 1

	ch <- randomRune
}

func terminate(done chan bool, c *int) {
	timer := time.NewTicker(time.Millisecond * 70)

	for tick := range timer.C {
		fmt.Printf("[%s] Checking c is %d\n", tick.Format(time.Kitchen), *c)
		if *c >= 2 {
			break
		}
	}

	fmt.Println("Terminating...")

	done <- true
	close(done)
}
