package main

import (
	"fmt"
	"time"
)

type RateLimiter struct {
	ticker time.Ticker
	bucket chan struct{}
	done   chan struct{}
}

func (c *RateLimiter) Wait() {
	<-c.bucket
}

func (c *RateLimiter) Close() {
	c.ticker.Stop()
	c.done <- struct{}{}
}

func main() {
	rl := getLimiter(10)

	for {
		select {
		case <-rl.bucket:
			fmt.Println("token received")
		}
	}
}

func getLimiter(n int) RateLimiter {
	ret := RateLimiter{
		ticker: *time.NewTicker(time.Millisecond * 300),
		bucket: make(chan struct{}, n),
		done:   make(chan struct{}),
	}

	for i := 0; i < n; i++ {
		ret.bucket <- struct{}{}
	}

	go func() {
		for {
			select {
			case <-ret.done:
				return
			case <-ret.ticker.C:
				ret.bucket <- struct{}{}
			}
		}
	}()

	go func() {
		time.Sleep(time.Second * 2)
		ret.Close()
	}()

	return ret
}
