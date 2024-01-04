package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type ProgressMeter struct {
	progress int64
}

func (p *ProgressMeter) Add() {
	atomic.AddInt64(&p.progress, 1)
}

func (p *ProgressMeter) Get() int64 {
	return atomic.LoadInt64(&p.progress)
}

func (p *ProgressMeter) Progress() {
	p.Add()
}

func orchestrator() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pm := ProgressMeter{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		process(ctx, &pm)
	}()
	go heartbeatsObserver(ctx, cancel, &pm)

	wg.Wait()
}

func heartbeatsObserver(ctx context.Context, cancel func(), pm *ProgressMeter) {
	var lastProgress int64
	ticker := time.NewTicker(time.Millisecond * 100)
	defer cancel()

	for tick := range ticker.C {

		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			p := pm.Get()
			if p-lastProgress > 100 {
				cancel()
				fmt.Println("Progress is not observed, cancelling")
				return
			}
			fmt.Println("Checking at ", tick.Format(time.TimeOnly))
		}
	}
}

func process(ctx context.Context, pm *ProgressMeter) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Process got cancel from context")
			return
		default:
		}
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(120)))

		pm.Progress()
	}
}
