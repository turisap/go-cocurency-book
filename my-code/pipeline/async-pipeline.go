package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"time"
)

func pipelineStage[IN any, OUT any](input <-chan IN, output chan<- OUT, process func(IN) OUT) {
	defer close(output)
	for data := range input {
		output <- process(data)
	}
}

func asyncPipeline(input *csv.Reader) {
	parseInput := make(chan []string)
	convertInput := make(chan Record)
	encodeInput := make(chan Record)
	outputCh := make(chan []byte)
	done := make(chan bool)

	go pipelineStage(parseInput, convertInput, parse)
	go pipelineStage(convertInput, encodeInput, convert)
	go pipelineStage(encodeInput, outputCh, encode)

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

		parseInput <- rec
	}

	<-done
}
