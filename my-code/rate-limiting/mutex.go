package main

import (
	"fmt"
	"sync"
	"time"
)

type Limiter struct {
	mu         sync.Mutex
	rate       int
	bucketSize int
	nTokens    int
	lastToken  time.Time
}

func NewLimiter(rate, limit int) *Limiter {
	return &Limiter{
		rate:       rate,
		bucketSize: limit,
		nTokens:    limit,
		lastToken:  time.Now(),
	}
}

func (l *Limiter) Wait() {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.nTokens > 0 {
		l.nTokens--
		return
	}

	tElapsed := time.Since(l.lastToken)
	period := time.Second / time.Duration(l.rate)
	nTokens := tElapsed.Nanoseconds() / period.Nanoseconds()
	l.nTokens = int(nTokens)

	if l.nTokens > l.bucketSize {
		l.nTokens = l.bucketSize
	}

	l.lastToken = l.lastToken.Add(time.Duration(nTokens) *
		period)

	if l.nTokens > 0 {
		l.nTokens--
		return
	}

	next := l.lastToken.Add(period)
	wait := next.Sub(time.Now())
	if wait >= 0 {
		time.Sleep(wait)
	}
	l.lastToken = next
}

func main() {
	limiter := NewLimiter(5, 8)

	for {
		fmt.Println("Making a request")
		limiter.Wait()
		fmt.Println("Got it back")
	}
}
