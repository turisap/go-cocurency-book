package main

import (
	"fmt"
	"sync"
)

func handleErrorsFromWorkerPool() {
	res1Ch := make(chan Result)
	res2Ch := make(chan Result)
	cancelCh := make(chan struct{})
	canceledCh := make(chan struct{})

	go w1(res1Ch, cancelCh, canceledCh)
	go w2(res2Ch, cancelCh, canceledCh)

	go func() {
		once := sync.Once{}
		for range cancelCh {
			once.Do(func() {
				close(canceledCh)
			})
		}
	}()

	res1 := <-res1Ch
	res2 := <-res2Ch

	fmt.Println("RES1", res1)

	fmt.Println("RES2", res2)
	fmt.Println("finish exec")
}

func w1(rCh chan Result, cancel chan struct{}, canceled chan struct{}) {
	res, err := computeResult(1)

	if err != nil {
		fmt.Println("Sending cancelling event from 1")
		cancel <- struct{}{}
		rCh <- Result{Error: err}
		return
	}

	select {
	case <-canceled:
		close(rCh)
		fmt.Println("Got canceling event in 1")
		return
	default:
		fmt.Println("Non blocking in 1")
	}

	rCh <- Result{Result: res}
}

func w2(rCh chan Result, cancel chan struct{}, canceled chan struct{}) {
	res, err := computeResult(1)

	if err != nil {
		fmt.Println("Sending cancelling event from 2")
		cancel <- struct{}{}
		rCh <- Result{Error: err}
		return
	}

	select {
	case <-canceled:
		close(rCh)
		fmt.Println("Got canceling event in 2")
		return
	default:
		fmt.Println("Non blocking in 2")
	}

	rCh <- Result{Result: res}
}
