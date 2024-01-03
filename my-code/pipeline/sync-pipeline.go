package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
)

type Record struct {
	Row    int     `json:"row"`
	Height float64 `json:"height"`
	Weight float64 `json:"weight"`
}

func newRecord(in []string) (rec Record, err error) {
	rec.Row, err = strconv.Atoi(in[0])
	if err != nil {
		return
	}
	rec.Height, err = strconv.ParseFloat(in[1], 64)
	if err != nil {
		return
	}
	rec.Weight, err = strconv.ParseFloat(in[2], 64)

	return
}

func parse(input []string) Record {
	rec, err := newRecord(input)
	if err != nil {
		panic(err)
	}
	return rec
}

func convert(input Record) Record {
	input.Height = 2.54 * input.Height
	input.Weight = 0.454 * input.Weight
	return input
}
func encode(input Record) []byte {
	data, err := json.Marshal(input)
	if err != nil {
		panic(err)
	}

	return data
}

func synchronousPipeline(input *csv.Reader) {
	// Skip the header row
	input.Read()
	for {
		rec, err := input.Read()
		if err == io.EOF {
			return
		}
		if err != nil {
			panic(err)
		}
		// The pipeline: parse, convert, encode
		out := encode(convert(parse(rec)))
		fmt.Println(string(out))
	}
}
