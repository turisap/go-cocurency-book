package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func contextFn() {
	c := context.Background()
	ctx1, cancel1 := context.WithCancel(c)

	go func(c context.Context) {
		ticker := time.NewTicker(time.Millisecond * 100)
		defer ticker.Stop()

		for {
			<-ticker.C
			fmt.Println("Tick")
		}
	}(ctx1)

	go func() {

		time.Sleep(time.Second * 3)
		fmt.Println("Cancelling context")
		cancel1()
	}()

	select {

	case <-ctx1.Done():
		fmt.Println("DONE")
	}

}

func withCancel() {
	c := context.Background()
	d := time.Now().Add(2 * time.Second)
	// decorator pattern
	ctx, cancel := context.WithDeadline(c, d)
	ticker := time.NewTicker(time.Millisecond * 200)
	defer ticker.Stop()
	defer cancel()

	for {
		<-ticker.C
		fmt.Println("Ticking")
		fmt.Println(ctx)

		if ctx.Err() == context.DeadlineExceeded {
			log.Fatal("ERROR", ctx.Err())
		}
	}
}
