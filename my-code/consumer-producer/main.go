package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	doneCh := make(chan bool)
	dataCh := make(chan int, 0)
	var producers sync.WaitGroup
	var consumers sync.WaitGroup

	for i := 0; i < 10; i++ {
		producers.Add(1)
		go producer(i, doneCh, dataCh, &producers)
	}
	for i := 0; i < 10; i++ {
		consumers.Add(1)
		go consumer(i, dataCh, &consumers)
	}

	time.Sleep(time.Second * 3)

	close(doneCh)
	producers.Wait()

	close(dataCh)
	consumers.Wait()
}

func producer(i int, done <-chan bool, data chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		val := rand.Int()

		time.Sleep(time.Millisecond *
			time.Duration(rand.Intn(1000)))

		select {
		case data <- val:
			fmt.Printf("Producer #%d send %v to the channel\n", i, val)
		case <-done:
			fmt.Printf("Shutting down producer #%d\n", i)
			return
		}
	}

}

func consumer(i int, data <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for val := range data {
		fmt.Printf("Consumer #%d consumed value %d\n", i, val)
	}
}
