package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	sign := make(chan os.Signal)

	signal.Notify(sign, syscall.SIGINT, syscall.SIGTERM)

	terminated := container(sign)

	<-terminated

	fmt.Println("Container cleaned up and terminated")
}

func container(sign chan os.Signal) <-chan bool {

	timer := time.NewTicker(time.Millisecond * 300)
	terminated := make(chan bool, 1)

outer:
	for {

		select {
		case tick := <-timer.C:
			fmt.Printf("[%s]: doing some container jobs\n", tick.Format(time.Kitchen))
		case <-sign:
			fmt.Println("Terminating due to signal")
			fmt.Println("Cleanup")
			terminated <- true

			close(terminated)
			close(sign)
			break outer
		}
	}

	return terminated
}
