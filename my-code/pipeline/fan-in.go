package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"sync"
	"time"
)

func workerPoolPipelineStage[IN any, OUT any](input <-chan IN,
	output chan<- OUT, process func(IN) OUT, numWorkers int) {
	// close output channel when all workers are done
	defer close(output)
	// Start the worker pool
	wg := sync.WaitGroup{}
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for data := range input {
				output <- process(data)
			}
		}()
	}
	// Wait for all workers to finish
	wg.Wait()
}

func asyncPipelineFan(input *csv.Reader) {
	parseInputCh := make(chan []string)
	convertInputCh := make(chan Record)
	encodeInputCh := make(chan Record)
	outputCh := make(chan []byte)
	done := make(chan bool)

	numWorkers := 2
	// Start pipeline stages and connect them
	go workerPoolPipelineStage(parseInputCh, convertInputCh, parse, numWorkers)
	go workerPoolPipelineStage(convertInputCh, encodeInputCh, convert, numWorkers)
	go workerPoolPipelineStage(encodeInputCh, outputCh, encode, numWorkers)

	go func() {
		for line := range outputCh {
			fmt.Println(string(line))
		}

		time.Sleep(time.Second * 3)
		done <- true
	}()

	input.Read()

	for {
		rec, err := input.Read()
		if err == io.EOF {
			return
		}
		if err != nil {
			panic(err)
		}

		parseInputCh <- rec
	}

	<-done
}
